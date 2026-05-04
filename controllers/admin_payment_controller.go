package controllers

import (
	"hris/config"
	"hris/models"
	"hris/utils"
	"hris/services"
	"time"

	"github.com/gin-gonic/gin"
)

func PaymentList(c *gin.Context) {

	var payments []models.Payment
	var accounts []models.Account

	config.DB.Where("is_owner = ?", false).Find(&accounts)

	// config.DB.Preload("Account").
	// 	Where("status = ?", "pending").
	// 	Find(&payments)

		config.DB.
		Preload("Account").
		Joins("JOIN subscriptions ON subscriptions.account_id = payments.account_id").
		Where("payments.status = ?", "pending").
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

	var payment models.Payment
	config.DB.First(&payment, id)

	// 🔥 update payment
	config.DB.Model(&payment).Update("status", "approved")

	// 🔥 aktifkan account
	config.DB.Model(&models.Account{}).
		Where("id = ?", payment.AccountID).
		Update("is_active", true)

	c.Redirect(302, "/owner/payments")
}

