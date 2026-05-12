package services

import (
	"hris/config"
	"hris/models"
	"time"
)

func CheckSubscriptions() {

	var accounts []models.Account
	config.DB.Where("is_owner = ?", false).Find(&accounts)

	now := time.Now()

	for _, acc := range accounts {

		// skip owner (biar gak ke lock)
		if acc.IsOwner {
			continue
		}

		// kalau belum pernah subscribe → skip
		// if acc.SubscriptionEnd == nil {
		// 	continue
		// }

		if acc.SubscriptionEnd.Before(now) {
			config.DB.Model(&acc).Update("islocked", true)
		} else {
			config.DB.Model(&acc).Update("islocked", false)
		}

		daysLeft := int(acc.SubscriptionEnd.Sub(now).Hours() / 24)

		// 🔥 H-7
		if daysLeft == 7 {
			SendReminder(acc, "7")
		}

		// 🔥 H-3
		if daysLeft == 3 {
			SendReminder(acc, "3")
		}

		// 🔥 EXPIRED
		if daysLeft <= 0 {
			config.DB.Model(&acc).Update("is_active", false)
		}
	}
}

func SendReminder(acc models.Account, day string) {

	message := "Subscription kamu akan berakhir dalam " + day + " hari."

	// 🔥 EMAIL
	SendEmailHTML(
		"client@email.com", // nanti ganti ambil dari DB
		"Subscription Reminder",
		message,
	)

	// 🔥 (optional) log ke DB
	config.DB.Model(&acc).Update("last_reminder_sent", time.Now())
}
