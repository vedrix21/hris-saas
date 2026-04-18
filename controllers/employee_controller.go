package controllers

import (
	"fmt"
	"hris/config"
	"hris/models"
	"hris/services"
	"hris/utils"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Employees(c *gin.Context) {
    var employees []models.Employee
    config.DB.Find(&employees)

    c.HTML(200, "employees.html", gin.H{
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