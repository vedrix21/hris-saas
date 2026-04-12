package controllers

import (
    "net/http"
    "hris/services"

    "github.com/gin-gonic/gin"
)

func ShowLogin(c *gin.Context) {
    accountCode := c.Query("account_code")

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

    c.Redirect(http.StatusFound, "/dashboard")
}

func Logout(c *gin.Context) {
    c.SetCookie("user", "", -1, "/", "", false, true)
    c.SetCookie("tenant", "", -1, "/", "", false, true)

    c.Redirect(302, "/login")
}