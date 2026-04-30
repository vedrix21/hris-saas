package models

import "time"

type Company struct {
    ID        uint
    CompanyID string `gorm:"unique"`
    AccountID uint

    Name      string
    Address   string

    CreatedAt time.Time
}