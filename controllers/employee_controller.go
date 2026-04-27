package controllers

import (
	"hris/config"
	"hris/models"
	"hris/services"
	"hris/utils"

	"github.com/gin-gonic/gin"
)

func Employees(c *gin.Context) {
	user := c.MustGet("user").(models.User)
	menus := services.GetSidebarByRole(user.Role)

	var employees []models.Employee
	config.DB.Find(&employees)

	utils.Render(c, []string{
		"templates/layout/base.html",
		"templates/layout/sidebar.html",
		"templates/admin/employee.html",
	}, gin.H{
		"Title":     "Employees",
		"Menus":     menus,
		"employees": employees,
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
