package controllers

import (
    "net/http"
    "hris/config"
    "hris/models"

    "github.com/gin-gonic/gin"
)

func ShowCreateAccount(c *gin.Context) {
    c.HTML(http.StatusOK, "create_account.html", nil)
}

func CreateAccount(c *gin.Context) {
    company := c.PostForm("company_name")
    code := c.PostForm("account_code")
    username := c.PostForm("username")
    password := c.PostForm("password")

    account := models.Account{
        Code:        code,
        CompanyName: company,
    }

    config.DB.Create(&account)

    user := models.User{
        Username:  username,
        Password:  password, // 🔥 sementara plain
        AccountID: account.ID,
    }

    config.DB.Create(&user)

    c.String(200, "Account berhasil dibuat")
}