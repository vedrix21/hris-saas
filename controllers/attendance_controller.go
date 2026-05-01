package controllers

import (
	"hris/config"
	"hris/models"
	"hris/services"
	"hris/utils"
	"time"

	"github.com/gin-gonic/gin"
)

func CheckIn(c *gin.Context) {
	user, _ := c.Cookie("user")

	attendance := models.Attendance{
		Username: user,
		CheckIn:  time.Now(),
	}

	config.DB.Create(&attendance)
	c.JSON(200, gin.H{"message": "Check-in success"})
}

func AttendancePage(c *gin.Context) {
	tenant, _ := c.Cookie("tenant")
	user := c.MustGet("user").(models.User)
	menus := services.GetSidebar(user.Role, tenant)

	var attendance []models.Attendance
	config.DB.Find(&attendance)

	utils.Render(c, []string{
		"templates/attendance/index.html",
	}, gin.H{
		"Title":      "Attendance",
		"Menus":      menus,
		"attendance": attendance,
		"Role":       user.Role,
	})
}
