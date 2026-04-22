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
    c.HTML(200, "owner/settings.html", gin.H{
    "title": "Settings",
    "Menus": menus,
    "CurrentPath": c.Request.URL.Path,
})
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
    db := config.DB

    var totalClients int64
    db.Model(&models.Account{}).
        Where("is_owner = ?", false).
        Count(&totalClients)

    user := c.MustGet("user").(models.User)

    menus := services.GetSidebarByRole(user.Role)

    c.HTML(200, "owner/dashboard.html", gin.H{
        "title": "Dashboard",
        "Menus": menus,
        "CurrentPath": c.Request.URL.Path,
        "accounts":      accounts,
        "success":       success,
        "total_clients": totalClients,
        
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

    c.Redirect(302, "/owner/dashboard")
}