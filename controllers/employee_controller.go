package controllers

import (
	"hris/config"
	"hris/models"
	"hris/services"
	"hris/utils"

	"github.com/gin-gonic/gin"
)

func Employees(c *gin.Context) {
	tenant, _ := c.Cookie("tenant")
	user := c.MustGet("user").(models.User)
	menus := services.GetSidebar(user.Role, tenant)

	var employees []models.Employee
	config.DB.Find(&employees)

	utils.Render(c, []string{
		"templates/employee/index.html",
	}, gin.H{
		"Title":       "Employees",
		"tenant":      tenant,
		"Menus":       menus,
		"employees":   employees,
		"Role":        user.Role,
		"CurrentPath": c.Request.URL.Path,
		"User":        user,
	})
}

func CreateEmployee(c *gin.Context) {
	emp := models.Employee{
		Name:  c.PostForm("name"),
		Email: c.PostForm("email"),
	}

	config.DB.Create(&emp)
	c.Redirect(302, "/employees")
}
