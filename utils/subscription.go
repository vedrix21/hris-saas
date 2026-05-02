package utils

import (
	"time"
	"hris/models"
)

func IsSubscriptionActive(acc models.Account) bool {
	if acc.SubscriptionEnd == nil {
		return false
	}

	return acc.SubscriptionEnd.After(time.Now())
}