package controllers

import (
	"hris/config"
	"hris/models"
	"hris/utils"
	"hris/services"

	"github.com/gin-gonic/gin"
)

func PaymentList(c *gin.Context) {

	var payments []models.Payment

	// 🔥 ambil payment + account
	config.DB.
		Preload("Account").
		Where("status = ?", "pending").
		Order("created_at desc").
		Find(&payments)

	// 🔥 ambil semua subscription pending
	var subs []models.Subscription
	config.DB.
		Where("status = ?", "pending").
		Find(&subs)

	// 🔥 mapping accountID -> subscription
	subMap := make(map[uint]models.Subscription)
	for _, s := range subs {
		subMap[s.AccountID] = s
	}

	

	tenant, _ := c.Cookie("tenant")
	user := c.MustGet("user").(models.User)
	menus := services.GetSidebar(user.Role, tenant)

	utils.Render(c, []string{
		"templates/owner/payments.html",
	}, gin.H{
		"title":       "Approve Payments",
		"Menus":       menus,
		"CurrentPath": c.Request.URL.Path,
		"Payments":    payments,
		"subs":        subMap,
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

func RejectPayment(c *gin.Context) {

	id := c.PostForm("id")
	note := c.PostForm("note")

	var payment models.Payment
	config.DB.First(&payment, id)

	config.DB.Model(&payment).Updates(map[string]interface{}{
		"status": "rejected",
		"note":   note,
	})

	// 🔥 simpan note ke subscription (opsional)
	config.DB.Model(&models.Subscription{}).
		Where("account_id = ? AND status = ?", payment.AccountID, "pending").
		Update("status", "rejected")

	// TODO: simpan note ke table lain kalau mau

	c.Redirect(302, "/owner/payments")
}


func GetPaymentProof(c *gin.Context) {
	id := c.Param("id")

	var payment models.Payment
	err := config.DB.First(&payment, id).Error
	if err != nil {
		c.JSON(404, gin.H{"error": "not found"})
		return
	}

	url, err := utils.GeneratePresignedURL(payment.Proof)
	if err != nil {
		c.JSON(500, gin.H{"error": "failed generate url"})
		return
	}

	c.JSON(200, gin.H{
		"url": url,
	})
}