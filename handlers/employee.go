
package handlers

import (
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "hris/models"
)

func GetEmployees(c *gin.Context) {
    db := c.MustGet("tenantDB").(*gorm.DB)
    var data []models.Employee
    db.Find(&data)
    c.JSON(200, data)
}

func CreateEmployee(c *gin.Context) {
    db := c.MustGet("tenantDB").(*gorm.DB)
    var e models.Employee
    c.BindJSON(&e)
    db.Create(&e)
    c.JSON(200, e)
}

func UpdateEmployee(c *gin.Context) {
    db := c.MustGet("tenantDB").(*gorm.DB)
    id := c.Param("id")

    var e models.Employee
    db.First(&e, id)
    c.BindJSON(&e)
    db.Save(&e)

    c.JSON(200, e)
}

func DeleteEmployee(c *gin.Context) {
    db := c.MustGet("tenantDB").(*gorm.DB)
    id := c.Param("id")
    db.Delete(&models.Employee{}, id)
    c.JSON(200, gin.H{"message":"deleted"})
}
