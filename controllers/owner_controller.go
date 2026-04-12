package controllers

import (
    "net/http"
    "hris/config"
    "hris/models"

    "github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func ShowCreateAccount(c *gin.Context) {
    c.HTML(http.StatusOK, "create_account.html", nil)
}

func ShowSettings(c *gin.Context) {
    c.HTML(200, "settings.html", nil)
}

func CreateAccount(c *gin.Context) {
    company := c.PostForm("company_name")
    code := c.PostForm("account_code")
    username := c.PostForm("username")
    password := c.PostForm("password")

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        c.String(500, "Failed to hash password")
        return
    }

    account := models.Account{
        Code:        code,
        CompanyName: company,
		Package:     "basic",
		UserLimit:   5,
		IsActive:    true,
    }

    config.DB.Create(&account)

    user := models.User{
        Username:  username,
        Password:  string(hashedPassword), // 🔥 sementara plain
        AccountID: account.ID,
    }

    config.DB.Create(&user)

    c.Redirect(302, "/owner/dashboard?success=1")
}

func OwnerDashboard(c *gin.Context) {
    var accounts []models.Account

    config.DB.Find(&accounts)

    success := c.Query("success")

    c.HTML(200, "owner_dashboard.html", gin.H{
        "accounts": accounts,
        "success":  success,
    })
}

func SaveSettings(c *gin.Context) {
    logo := c.PostForm("logo_url")
    color := c.PostForm("theme_color")

    tenant, _ := c.Cookie("tenant")

    config.DB.Model(&models.Account{}).
        Where("code = ?", tenant).
        Updates(map[string]interface{}{
            "logo_url": logo,
            "theme_color": color,
        })

    c.Redirect(302, "/dashboard")
}