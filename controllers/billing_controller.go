package controllers

import (
	
	"hris/config"
	"hris/models"
    "time"
	
	

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
		"package":             plan.PlanName,
		"monthly_fee":         plan.Price,
		"user_limit":          plan.Limituser,
		"subscription_start":  now,
		"subscription_end":    end,
		"is_active":           false, // 🔥 belum aktif sampai bayar
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