package controllers

import (
	
	"hris/config"
	"hris/models"

	
	
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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
    var data []models.Attendance
    config.DB.Order("created_at desc").Find(&data)

    c.HTML(200, "attendance.html", gin.H{
        "data": data,
    })
}

