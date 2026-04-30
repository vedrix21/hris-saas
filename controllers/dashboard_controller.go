package controllers

import (
    "github.com/gin-gonic/gin"
    "hris/utils"
    "hris/config"
    "hris/models"
    "hris/services"
    
)

func Dashboard(c *gin.Context) {
    tenant, _ := c.Cookie("tenant")

    var totalEmployee int64
    config.DB.Model(&models.Employee{}).Count(&totalEmployee)


	user := c.MustGet("user").(models.User)

	menus := services.GetSidebarByRole(user.Role)

    utils.Render(c, []string{
        "templates/layout/base.html",
        "templates/layout/sidebar.html",
        "templates/components/loading.html",
        "templates/admin/dashboard.html",
    }, gin.H{
        "title": "Dashboard",
        "tenant": tenant,
        "totalEmployee": totalEmployee,
        "Menus":         menus,
        "CurrentPath":   c.Request.URL.Path,
    })
}

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