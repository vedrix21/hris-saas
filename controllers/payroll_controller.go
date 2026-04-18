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
	"golang.org/x/crypto/bcrypt"
)

func RunPayroll(c *gin.Context) {
    err := services.ProcessPayroll()
    if err != nil {
        c.String(500, "Payroll failed")
        return
    }

    c.String(200, "Payroll success")
}

func PayrollPage(c *gin.Context) {
    var payrolls []models.Payroll
    config.DB.Find(&payrolls)

    c.HTML(200, "payroll.html", gin.H{
        "payrolls": payrolls,
    })
}