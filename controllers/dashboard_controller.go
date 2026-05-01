package controllers

import (
	"hris/config"
	"hris/models"
	"hris/services"
	"hris/services/modules"
	"hris/utils"

	"github.com/gin-gonic/gin"
)

func Dashboard(c *gin.Context) {

	var totalEmployee int64
	config.DB.Model(&models.Employee{}).Count(&totalEmployee)
	tenant, _ := c.Cookie("tenant")
	user := c.MustGet("user").(models.User)
	menus := services.GetSidebar(user.Role, tenant)

	utils.Render(c, []string{
		"templates/layout/base.html",
		"templates/layout/sidebar.html",
		"templates/components/loading.html",
		"templates/admin/dashboard.html",
	}, gin.H{
		"title":         "Dashboard",
		"tenant":        tenant,
		"totalEmployee": totalEmployee,
		"Menus":         menus,
		"CurrentPath":   c.Request.URL.Path,
		"User":          user,
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
