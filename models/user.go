package models

import "time"

type User struct {
    ID        uint
    Username  string `gorm:"type:varchar(50);index"`
    Password  string
    AccountID uint
	Role	  string
	ResetToken     string
	ResetTokenExp *time.Time
	Email          string `gorm:"type:varchar(100);index"`
	FullName       string
}