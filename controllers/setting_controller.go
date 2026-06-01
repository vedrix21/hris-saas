package controllers

import (
	
	"hris/models"
	"hris/services"
	"hris/utils"
	"github.com/gin-gonic/gin"
)


func MigrationPage(c *gin.Context) {
	tenant, _ := c.Cookie("tenant")
	user := c.MustGet("user").(models.User)

	menus := services.GetSidebar(user.Role, tenant)

	utils.Render(c, []string{
		"templates/settings/data_migration.html",
	}, gin.H{
		"Title":       "Data Migration",
		"tenant":      tenant,
		"Menus":       menus,
		"Role":        user.Role,
		"CurrentPath": c.Request.URL.Path,
		"User":        user,
	})
}