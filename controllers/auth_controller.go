package controllers

import (
	"fmt"
	"hris/config"
	"hris/models"
	"hris/services"
	"hris/utils"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	// "golang.org/x/crypto/bcrypt"
)

func Home(c *gin.Context) {

	user, err := c.Cookie("user")
	role, _ := c.Cookie("role")

	if err != nil || user == "" {
		c.Redirect(302, "/login")
		return
	}

	if role == "owner" {
		c.Redirect(302, "/owner/dashboard")
		return
	}

	c.Redirect(302, "/dashboard")
}

func ShowLogin(c *gin.Context) {
	// user, err := c.Cookie("user")
	// fmt.Println("Masuk ke show login")

	// if err == nil && user != "" {
	// 	role, _ := c.Cookie("role")

	// 	if role == "owner" {
	// 		c.Redirect(302, "/owner/dashboard")
	// 	} else {
	// 		c.Redirect(302, "/dashboard")
	// 	}
	// 	return
	// }

	renderLogin(c, "")
}

func Login(c *gin.Context) {
	// accountCode := strings.ToLower(c.PostForm("account_code"))
	accountCode := c.PostForm("account_code")
	username := c.PostForm("username")
	password := c.PostForm("password")

	if accountCode == "" || username == "" || password == "" {
		renderLogin(c, "Semua field wajib diisi")
		return
	}

	user, account, err := services.Login(accountCode, username, password)
	if err != nil {
		// fmt.Println("LOGIN ERROR:", err.Error()) // 🔥 log ke terminal
		// renderLogin(c, err.Error())              // 🔥 tampilkan ke UI
		renderLogin(c, "Invalid Credential")
		return
	}

	domain := "" // default

	if os.Getenv("ENV") == "production" {
		domain = "app.aitherhr.com"
	}

	c.SetCookie("tenant", account.Code, 3600, "/", domain, true, true)
	c.SetCookie("user", user.Username, 3600, "/", domain, true, true)
	c.SetCookie("role", user.Role, 3600, "/", domain, true, true)

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
	c.SetCookie("role", "", -1, "/", "", false, true)

	c.Redirect(302, "/login")
}

func ShowForgotPassword(c *gin.Context) {
	utils.RenderAuth(c, []string{
		"templates/forgot_password.html",
	}, gin.H{})
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

	baseURL := os.Getenv("BASE_URL")
	if baseURL == "" {
		baseURL = "http://localhost:8080"
	}

	resetLink := baseURL + "/reset-password?token=" + token

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

	c.String(200, "Link reset password telah dikirim ke email")
}

func ShowResetPassword(c *gin.Context) {
	token := c.Query("token")
	utils.Render(c, []string{
		"templates/reset_password.html",
	}, gin.H{
		"token": token,
	})
}

func ResetPassword(c *gin.Context) {
	token := c.PostForm("token")
	newPassword := c.PostForm("password")

	var user models.User
	if err := config.DB.Where("TRIM(reset_token) = ?", token).First(&user).Error; err != nil {
		fmt.Println("DB ERROR:", err)
		c.String(400, "Invalid token")
		return
	}

	if user.ResetTokenExp == nil || time.Now().After(*user.ResetTokenExp) {
		c.String(400, "Token expired")
		return
	}

	// hashed, _ := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	// user.Password = string(hashed)
	user.Password = string(newPassword)

	user.ResetToken = ""
	user.ResetTokenExp = nil

	if err := config.DB.Save(&user).Error; err != nil {
		fmt.Println("SAVE ERROR:", err)
		c.String(500, "Failed update password")
		return
	}

	renderLogin(c, "Password berhasil direset, silakan login")

	// utils.Render(c, []string{
	// 	"templates/login.html",
	// }, gin.H{
	// 	"success": "Password berhasil direset, silakan login",
	// })
}

func renderLogin(c *gin.Context, errorMsg string) {
	utils.RenderAuth(c, []string{
		"templates/auth/login.html",
	}, gin.H{
		"error": errorMsg,
		"logo":  "/static/logo.png",
	})
}

func SwitchEnv(c *gin.Context) {
	env := c.Param("env") // dev / prod

	c.SetCookie("env", env, 3600, "/", "", false, true)

	c.Redirect(302, "/dashboard")
}
