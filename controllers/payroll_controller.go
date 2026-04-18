package controllers

import (
	
	"hris/config"
	"hris/models"
	"hris/services/modules"
	
	"github.com/gin-gonic/gin"
)

func RunPayroll(c *gin.Context) {

    account, _ := c.Cookie("tenant")
    env := utils.GetEnv(c)

    err := modules.ProcessPayroll(account, env)

    if err != nil {
        c.String(500, "Payroll failed: "+err.Error())
        return
    }

    c.String(200, "Payroll success 🚀")
}

func PayrollPage(c *gin.Context) {
    var payrolls []models.Payroll
    config.DB.Find(&payrolls)

    c.HTML(200, "payroll.html", gin.H{
        "payrolls": payrolls,
    })
}