package controllers

import (
	"hris/config"
	"hris/models"
	"hris/utils"
	"hris/services"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func PaymentList(c *gin.Context) {

	var payments []models.Payment
	var accounts []models.Account

	config.DB.Where("is_owner = ?", false).Find(&accounts)

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

	filename := fmt.Sprintf("%d_%s", time.Now().Unix(), file.Filename)
	path := "static/uploads/" + filename

	if err := c.SaveUploadedFile(file, path); err != nil {
		c.JSON(500, gin.H{"error": "upload failed"})
		return
	}

	tenant, _ := c.Cookie("tenant")

	var acc models.Account
	config.DB.Where("code = ?", tenant).First(&acc)

	// 🔥 ambil subscription pending
	// var sub models.Subscription
	// config.DB.Where("account_id = ? AND status = ?", acc.ID, "pending").
	// 	Order("created_at desc").
	// 	First(&sub)

	// 🔥 INSERT ke payments table (INI YANG KAMU BELUM ADA)
	payment := models.Payment{
		AccountID: acc.ID,
		Amount:    acc.MonthlyFee,
		Proof:     filename,
		Status:    "pending",
	}

	config.DB.Create(&payment)

	// 🔥 update subscription terakhir
	// config.DB.Model(&models.Subscription{}).
	// 	Where("account_id = ? AND status = ?", acc.ID, "pending").
	// 	Update("proof", filename)

	c.Redirect(302, "/billing")
}