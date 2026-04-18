package controllers

import (
    "net/http"
    "hris/config"
    "hris/models"
    "hris/services"

    "github.com/gin-gonic/gin"
	
)

func ShowCreateAccount(c *gin.Context) {
    c.HTML(http.StatusOK, "create_account.html", nil)
}

func ShowSettings(c *gin.Context) {
    c.HTML(200, "settings.html", nil)
}

func CreateAccount(c *gin.Context) {
    companyName := c.PostForm("company_name")

    account, username, password, err := services.CreateTenant(companyName)
    if err != nil {
        c.String(500, "Failed create account")
        return
    }

    c.JSON(200, gin.H{
        "account_code": account.Code,
        "username":     username,
        "password":     password,
    })
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