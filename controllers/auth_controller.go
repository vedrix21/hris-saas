package controllers

import (
	"fmt"
    "net/http"
	"time"
    "os"
    "hris/services"
	"hris/config"
    "hris/models"
	"hris/utils"

    "github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func ShowLogin(c *gin.Context) {
    accountCode := c.PostForm("account_code")

    var account models.Account
    config.DB.Where("code = ?", accountCode).First(&account)

    logo := account.LogoURL
    if logo == "" {
        logo = "/static/logo.png"
    }

    color := account.ThemeColor
    if color == "" {
        color = "#4F46E5"
    }

    c.HTML(200, "login.html", gin.H{
        "logo":  logo,
        "color": color,
    })
}

func Login(c *gin.Context) {
    accountCode := c.PostForm("account_code")
    username := c.PostForm("username")
    password := c.PostForm("password")

    user, account, err := services.Login(accountCode, username, password)
    if err != nil {
        c.HTML(http.StatusUnauthorized, "login.html", gin.H{
            "error": "Login gagal",
        })
        return
    }

    // set cookie tenant
    c.SetCookie("tenant", account.Code, 3600, "/", "", false, true)

    // set session user (simple)
    c.SetCookie("user", user.Username, 3600, "/", "", false, true)

	// set cookie role
	c.SetCookie("role", user.Role, 3600, "/", "", false, true)

    switch user.Role {
	case "owner":
		c.Redirect(http.StatusFound, "/owner/dashboard")
	case "admin":
		c.Redirect(http.StatusFound, "/dashboard")
	default:
		c.Redirect(http.StatusFound, "/dashboard")
	}
}

func Logout(c *gin.Context) {
    c.SetCookie("user", "", -1, "/", "", false, true)
    c.SetCookie("tenant", "", -1, "/", "", false, true)

    c.Redirect(302, "/login")
}

func ShowForgotPassword(c *gin.Context) {
    c.HTML(200, "forgot_password.html", nil)
}

func ForgotPassword(c *gin.Context) {

    email := c.PostForm("email")

    var user models.User
    if err := config.DB.Where("email = ?", email).First(&user).Error; err != nil {
        c.String(404, "Email tidak ditemukan")
        return
    }

    token := utils.GenerateToken()

    config.DB.Model(&user).Updates(map[string]interface{}{
        "reset_token":     token,
        "reset_token_exp": time.Now().Add(1 * time.Hour),
    })

    resetLink := "https://app.hrflowapp.com/reset-password?token=" + token

    body := `
    <h3>Reset Password</h3>
    <p>Klik link berikut untuk reset password:</p>
    <a href="` + resetLink + `">Reset Password</a>
    `

    err := services.SendEmailHTML(user.Email, "Reset Password AitherHR", body)
	if err != nil {
        fmt.Println("EMAIL ERROR:", err)
        fmt.Println("SMTP HOST:", os.Getenv("SMTP_HOST"))
        fmt.Println("SMTP PORT:", os.Getenv("SMTP_PORT"))
    } else {
        fmt.Println("EMAIL SENT SUCCESS")
    }

    //c.String(200, "Link reset password telah dikirim ke email")
}

func ShowResetPassword(c *gin.Context) {
    token := c.Query("token")
    c.HTML(200, "reset_password.html", gin.H{"token": token})
}

func ResetPassword(c *gin.Context) {
    token := c.PostForm("token")
    newPassword := c.PostForm("password")

    var user models.User
    if err := config.DB.Where("reset_token = ?", token).First(&user).Error; err != nil {
        fmt.Println("DB ERROR:", err)
        c.String(400, "Invalid token")
        return
    }

    if user.ResetTokenExp == nil || time.Now().After(*user.ResetTokenExp) {
        c.String(400, "Token expired")
        return
    }

    hashed, _ := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
    user.Password = string(hashed)

    user.ResetToken = ""
    user.ResetTokenExp = nil

    if err := config.DB.Save(&user).Error; err != nil {
        fmt.Println("SAVE ERROR:", err)
        c.String(500, "Failed update password")
        return
    }

    c.String(200, "Password berhasil direset")
}