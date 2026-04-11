package controllers

import (
    "net/http"
    "hris/services"

    "github.com/gin-gonic/gin"
)

func ShowLogin(c *gin.Context) {
    c.HTML(http.StatusOK, "login.html", nil)
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

    c.Redirect(http.StatusFound, "/dashboard")
}
func Dashboard(c *gin.Context) {
    tenant, _ := c.Cookie("tenant")

    c.HTML(200, "dashboard.html", gin.H{
        "tenant": tenant,
    })
}
func Logout(c *gin.Context) {
    c.SetCookie("user", "", -1, "/", "", false, true)
    c.SetCookie("tenant", "", -1, "/", "", false, true)

    c.Redirect(302, "/login")
}