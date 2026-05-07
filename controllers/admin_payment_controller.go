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
		"payments":    payments,
		"subs":        subMap,
	})
}

func ApprovePayment(c *gin.Context) {

	id := c.PostForm("id")

	var payment models.Payment
	if err := config.DB.First(&payment, id).Error; err != nil {
		c.String(404, "payment not found")
		return
	}

	// 🔥 ambil account
	var acc models.Account
	if err := config.DB.First(&acc, payment.AccountID).Error; err != nil {
		c.String(404, "account not found")
		return
	}

	// 🔥 hitung tanggal subscription baru
	now := time.Now()

	var newEnd time.Time

	// kalau masih aktif → extend dari end lama
	if acc.SubscriptionEnd.After(now) {
		newEnd = acc.SubscriptionEnd.AddDate(0, 1, 0)
	} else {
		// kalau sudah expired → mulai dari sekarang
		newEnd = now.AddDate(0, 1, 0)
	}

	// 🔥 update payment
	config.DB.Model(&payment).Update("status", "approved")

	// 🔥 update account
	config.DB.Model(&acc).Updates(map[string]interface{}{
		"is_active":        true,
		"is_locked":        false,
		"subscription_end": newEnd,
	})

	// 🔥 update subscription terakhir (optional tapi recommended)
	config.DB.Model(&models.Subscription{}).
		Where("account_id = ? AND status = ?", acc.ID, "pending").
		Update("status", "approved")

	// 🔥 kirim email
	// 🔥 Email content
	body := `
	Halo ` + acc.PicName + `,
	<br>
	<br>
	<p>Pembayaran Anda telah berhasil diverifikasi.</p>

	<p>Akun Anda sudah aktif kembali dan berlaku sampai: ` + newEnd.Format("02 Jan 2006") + `</p>
	<p>Terima kasih telah menggunakan AitherHR 🚀 </p>

	<br>
	<a href="https://app.aitherhr.com">Login AitherHR</a>
	<br>
	<br>
	Regards,
	AitherHR Team
	`

	// 🔥 kirim ke email kamu
	err := services.SendEmailHTML(
		acc.PicEmail,
		"[AitherHR] Verifikasi Pembayaran Berhasil",
		body,
	)
	if err != nil {
		c.String(500, "Approval successful but email failed")
		return
	}

	c.Redirect(302, "/owner/payments")
}

func RejectPayment(c *gin.Context) {

	id := c.PostForm("id")
	note := c.PostForm("note")

	var payment models.Payment
	if err := config.DB.First(&payment, id).Error; err != nil {
		c.String(404, "payment not found")
		return
	}

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