package utils

import (
	"hris/models"
	"time"
)

func IsSubscriptionActive(acc models.Account) bool {
	if acc.SubscriptionEnd == nil {
		return false
	}

	return acc.SubscriptionEnd.After(time.Now())
}
func IsAccountLocked(acc models.Account) bool {

	if acc.SubscriptionEnd == nil {
		return true
	}

	if !acc.IsActive {
		return true
	}

	return time.Now().After(*acc.SubscriptionEnd)
}

func DaysLeft(acc models.Account) int {
	if acc.SubscriptionEnd == nil {
		return 0
	}
	return int(acc.SubscriptionEnd.Sub(time.Now()).Hours() / 24)
}
