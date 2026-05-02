package controllers

import (
	"hris/config"
	"hris/models"
	"hris/utils"
	"hris/services"

	"github.com/gin-gonic/gin"
)

func PaymentList(c *gin.Context) {

	var payments []models.Subscription
	var accounts []models.Account

	config.DB.Preload("Account").
		Where("status = ?", "pending").
		Find(&payments)

	tenant, _ := c.Cookie("tenant")
	user := c.MustGet("user").(models.User)
	menus := services.GetSidebar(user.Role, tenant)

	utils.Render(c, []string{
		"templates/owner/payments.html",
	}, gin.H{
		"title":    "Approve Payments",
		"Menus":         menus,
		"CurrentPath":   c.Request.URL.Path,
		"accounts":      accounts,
		"payments": payments,
	})
}

func ApprovePayment(c *gin.Context) {

	id := c.PostForm("id")

	var sub models.Subscription
	config.DB.First(&sub, id)

	// 🔥 update subscription
	config.DB.Model(&sub).Update("status", "active")

	// 🔥 aktifkan account
	config.DB.Model(&models.Account{}).
		Where("id = ?", sub.AccountID).
		Update("is_active", true)

	c.Redirect(302, "/owner/payments")
}

func UploadPayment(c *gin.Context) {

	file, err := c.FormFile("proof")
	if err != nil {
		c.JSON(400, gin.H{"error": "file required"})
		return
	}

	filename := file.Filename
	c.SaveUploadedFile(file, "static/uploads/"+filename)

	tenant, _ := c.Cookie("tenant")

	var acc models.Account
	config.DB.Where("code = ?", tenant).First(&acc)

	// 🔥 update subscription terakhir
	config.DB.Model(&models.Subscription{}).
		Where("account_id = ? AND status = ?", acc.ID, "pending").
		Update("proof", filename)

	c.Redirect(302, "/billing")
}