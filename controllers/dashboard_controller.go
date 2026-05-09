package controllers

import (
	"hris/config"
	"hris/models"
	"hris/services"
	"hris/services/modules"
	"hris/utils"
	"time"

	"github.com/gin-gonic/gin"
)

func Dashboard(c *gin.Context) {

	var totalEmployee int64
	config.DB.Model(&models.Employee{}).Count(&totalEmployee)
	tenant, _ := c.Cookie("tenant")
	user := c.MustGet("user").(models.User)
	menus := services.GetSidebar(user.Role, tenant)

	var account models.Account
	config.DB.Where("code = ?", tenant).First(&account)
	daysLeft := int(account.SubscriptionEnd.Sub(time.Now()).Hours() / 24)

	warning := false
	if daysLeft <= 7 {
		warning = true
	}

	utils.Render(c, []string{
		"templates/admin/dashboard.html",
	}, gin.H{
		"title":               "Dashboard",
		"tenant":              tenant,
		"totalEmployee":       totalEmployee,
		"Menus":               menus,
		"CurrentPath":         c.Request.URL.Path,
		"User":                user,
		"SubscriptionWarning": warning,
		"DaysLeft":            daysLeft,
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
