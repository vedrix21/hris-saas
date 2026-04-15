package controllers

import (
    "github.com/gin-gonic/gin"
    "hris/services/modules"
    "hris/utils"
)

func Dashboard(c *gin.Context) {
    tenant, _ := c.Cookie("tenant")

    c.HTML(200, "dashboard.html", gin.H{
        "tenant": tenant,
        "totalEmployee": 120, // 🔥 nanti diganti dari DB
    })
}

import (
    "hris/services/modules"
    "hris/utils"
)

func RunProcess(c *gin.Context) {

    account, _ := c.Cookie("tenant")
    env := utils.GetEnv(c)

    // 🔥 isolasi error per module
    payrollErr := modules.ProcessPayroll(account, env)

    var message string

    if payrollErr != nil {
        message += "Payroll error: " + payrollErr.Error() + "\n"
    } else {
        message += "Payroll success\n"
    }

    c.String(200, message)
}