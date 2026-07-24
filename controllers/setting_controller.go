package controllers

import (
	"fmt"
	"hris/models"
	"hris/services"
	"hris/utils"
	"hris/config"
	"time"
	"github.com/gin-gonic/gin"
)


func MigrationPage(c *gin.Context) {
	tenant, _ := c.Cookie("tenant")
	user := c.MustGet("user").(models.User)

	menus := services.GetSidebar(user.Role, tenant)

	var account models.Account
	config.DB.Where("code = ?", tenant).First(&account)
	daysLeft := int(account.SubscriptionEnd.Sub(time.Now()).Hours() / 24)

	var userdb models.User
	err := config.DB.Where("account_id = ? and username = ?", account.ID, user.Username).First(&userdb)
	if err != nil {
		fmt.Printf("Error fetching user from DB: %v\n", err)
	}

	warning := false
	if daysLeft <= 7 {
		warning = true
	}

	utils.Render(c, []string{
		"templates/settings/data_migration.html",
	}, gin.H{
		"Title":       "Data Migration",
		"tenant":      tenant,
		"Menus":       menus,
		"Role":        user.Role,
		"CurrentPath": c.Request.URL.Path,
		"User":        user,
		"SubscriptionWarning": warning,
		"DaysLeft":            daysLeft,
		"UserDB":              userdb,
	})
}

func DownloadMigrationTemplate(c *gin.Context) {
	
}
func ImportEmployeeData(c *gin.Context) {
    if _, err := c.FormFile("file"); err != nil {
        c.JSON(400, gin.H{"error": "File upload failed"})
        return
    }

    c.JSON(200, gin.H{"message": "Import employee belum diimplementasikan"})
}

func ImportLeaveData(c *gin.Context) {
    if _, err := c.FormFile("file"); err != nil {
        c.JSON(400, gin.H{"error": "File upload failed"})
        return
    }

    c.JSON(200, gin.H{"message": "Import leave belum diimplementasikan"})
}

func ImportPositionData(c *gin.Context) {
    if _, err := c.FormFile("file"); err != nil {
        c.JSON(400, gin.H{"error": "File upload failed"})
        return
    }

    c.JSON(200, gin.H{"message": "Import position belum diimplementasikan"})
}

func ImportOrgUnitData(c *gin.Context) {
    if _, err := c.FormFile("file"); err != nil {
        c.JSON(400, gin.H{"error": "File upload failed"})
        return
    }

    c.JSON(200, gin.H{"message": "Import org unit belum diimplementasikan"})
}