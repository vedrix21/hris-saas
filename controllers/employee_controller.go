package controllers

import (
	"hris/config"
	"hris/models"
	"hris/services"
	"hris/utils"
	"time"

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

	joinDate, _ := time.Parse("2006-01-02", c.PostForm("join_date"))

	employee := models.Employee{
		FirstName:        c.PostForm("first_name"),
		MiddleName:       c.PostForm("middle_name"),
		LastName:         c.PostForm("last_name"),
		BirthPlace:       c.PostForm("birth_place"),
		BirthDate:        c.PostForm("birth_date"),
		Gender:           c.PostForm("gender"),
		Religion:         c.PostForm("religion"),
		Email:            c.PostForm("email"),
		Phone:            c.PostForm("phone"),
		Address:          c.PostForm("address"),
		EmployeeCode:     c.PostForm("employee_code"),
		Position:         c.PostForm("position"),
		Department:       c.PostForm("department"),
		JoinDate:         joinDate,
		EmploymentStatus: c.PostForm("employment_status"),
	}

	config.DB.Create(&employee)
	c.Redirect(302, "/employees")
}
