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

    joinDate, _ := time.Parse(
        "2006-01-02",
        c.PostForm("join_date"),
    )

    firstName := c.PostForm("first_name")
    middleName := c.PostForm("middle_name")
    lastName := c.PostForm("last_name")

    fullName := firstName

    if middleName != "" {
        fullName += " " + middleName
    }

    if lastName != "" {
        fullName += " " + lastName
    }

    employee := models.Employee{
        FirstName:        firstName,
        MiddleName:       middleName,
        LastName:         lastName,
        FullName:         fullName,
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
        EmployeeStatus:   "Active",
        CreatedAt:        time.Now().Unix(),
    }

    config.DB.Create(&employee)

    c.Redirect(302, "/employees")
}


func DeleteEmployee(c *gin.Context) {

    id := c.Param("id")

    config.DB.Delete(&models.Employee{}, id)

    c.Redirect(302, "/employees")
}

func UpdateEmployee(c *gin.Context) {

    id := c.Param("id")

    var employee models.Employee

    config.DB.First(&employee, id)

    firstName := c.PostForm("first_name")
    middleName := c.PostForm("middle_name")
    lastName := c.PostForm("last_name")

    fullName := firstName

    if middleName != "" {
        fullName += " " + middleName
    }

    if lastName != "" {
        fullName += " " + lastName
    }

    joinDate, _ := time.Parse(
        "2006-01-02",
        c.PostForm("join_date"),
    )

    employee.FirstName = firstName
    employee.MiddleName = middleName
    employee.LastName = lastName
    employee.FullName = fullName
    employee.Email = c.PostForm("email")
    employee.Phone = c.PostForm("phone")
    employee.Position = c.PostForm("position")
    employee.Department = c.PostForm("department")
    employee.JoinDate = joinDate

    config.DB.Save(&employee)

    c.Redirect(302, "/employees")
}