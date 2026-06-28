package controllers

import (
	"errors"
	"fmt"
	"hris/config"
	"hris/models"
	"hris/utils"
	"time"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func UpgradePlan(c *gin.Context) {
	accountCode, _ := c.Cookie("tenant")

	var acc models.Account
	config.DB.Where("code = ?", accountCode).First(&acc)

	// 🔥 ambil plan dari form
	planID := c.PostForm("plan_id")

	var plan models.Subscriptionplan
	config.DB.First(&plan, planID)

	// 🔥 simpan plan lama sebelum diubah
	oldPlan := acc.Package

	// 🔥 update account
	now := time.Now()
	end := now.AddDate(0, 1, 0) // 1 bulan

	config.DB.Model(&acc).Updates(map[string]interface{}{
		"package":            plan.PlanName,
		"monthly_fee":        plan.Price,
		"user_limit":         plan.Limituser,
		"subscription_start": now,
		"subscription_end":   end,
		"is_active":          false, // 🔥 belum aktif sampai bayar
	})

	// 🔥 simpan history subscription
	config.DB.Create(&models.Subscription{
		AccountID: acc.ID,
		FromPlan:  oldPlan,
		ToPlan:    plan.PlanName,
		Status:    "pending", // 🔥 belum bayar
	})

	c.Redirect(302, "/billing")
}

func UploadPayment(c *gin.Context) {

	fmt.Println("UPLOAD HIT")

	tenant, _ := c.Cookie("tenant")

	var acc models.Account
	config.DB.Where("code = ?", tenant).First(&acc)

	// 🔥 ambil subscription pending
	// var sub models.Subscription
	// config.DB.Where("account_id = ? AND status = ?", acc.ID, "pending").
	// 	Order("created_at desc").
	// 	First(&sub)

	var existing models.Payment

	err := config.DB.
		Where("account_id = ? AND status = ?", acc.ID, "pending").
		First(&existing).Error

	if err == nil {
		// 🔥 masih ada pending payment
		c.JSON(400, gin.H{
			"error": "Masih ada pembayaran yang sedang direview. Harap tunggu approval.",
		})
		return
	}

	file, err := c.FormFile("proof")
	if err != nil {
		c.JSON(400, gin.H{"error": "file required"})
		return
	}

	// 🔥 generate path
	// path, err := utils.SaveUpload(file, "payments")
	// if err != nil {
	//     c.JSON(400, gin.H{"error": err.Error()})
	//     return
	// }

	// // 💾 simpan file
	// if err := c.SaveUploadedFile(file, path); err != nil {
	//     c.JSON(500, gin.H{"error": "failed to save file"})
	//     return
	// }

	// 🌐 URL yang disimpan ke DB
	// url := "/" + path // penting!

	src, err := file.Open()
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to read file"})
		return
	}
	defer src.Close()

	// 🔥 upload ke S3 (Railway Bucket)
	key, err := utils.UploadToS3(src, file.Filename, file.Header.Get("Content-Type"))
	if err != nil {
		fmt.Println("S3 ERROR:", err)
		c.JSON(500, gin.H{"error": "failed to upload file"})
		return
	}

	payment := models.Payment{
		AccountID: acc.ID,
		Amount:    acc.MonthlyFee,
		Proof:     key, // 🔥 simpan key S3, bukan path lokal
		Status:    "pending",
	}

	result := config.DB.Create(&payment)

	if result.Error != nil {
		fmt.Println("DB ERROR:", result.Error)
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}

	// c.JSON(200, gin.H{
	//     "message": "upload success",
	//     "url":     url,
	// })

	subscription := models.Subscription{
		AccountID: acc.ID,
		FromPlan:  acc.Package,
		ToPlan:    acc.Package,
		Status:    "pending",
	}

	result = config.DB.Create(&subscription)

	if result.Error != nil {
		fmt.Println("DB ERROR:", result.Error)
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}

	fmt.Println("INSERT SUCCESS, ID:", subscription.ID)

	// 🔥 update subscription terakhir
	// config.DB.Model(&models.Subscription{}).
	// 	Where("account_id = ? AND status = ?", acc.ID, "pending").
	// 	Update("proof", filename)

	c.Redirect(302, "/billing")
}

func BillingPage(c *gin.Context) {
	accountCode, _ := c.Cookie("tenant")
	
	user := c.MustGet("user").(models.User)
	

	var acc models.Account
	config.DB.Where("code = ?", accountCode).First(&acc)

	var payment models.Payment
	err := config.DB.Where("account_id = ?", acc.ID).
		Order("created_at desc").
		First(&payment).Error

	if err != nil {
		// ✅ abaikan kalau memang belum ada payment
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}
	}

	islocked := utils.IsAccountLocked(acc)

	if !islocked {
		c.Redirect(302, "/dashboard")
		return
	}

	utils.Render(c, []string{
		"templates/billing.html",
	}, gin.H{
		"account":  acc, // ✅ ini harus struct
		"payment":  payment,
		"isLocked": islocked,
		"daysLeft": utils.DaysLeft(acc),
		"User":        user,
	})
}
