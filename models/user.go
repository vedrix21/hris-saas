package models

import "time"

type User struct {
    ID        uint
    Username  string
    Password  string
    AccountID uint
	Role	  string
	ResetToken     string
	ResetTokenExp *time.Time
	Email          string
}