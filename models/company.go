package models

import "time"

type Company struct {
    ID        uint
    CompanyID string `gorm:"unique"`
    AccountCode string // e.g. "TENANT001"
    AccountID uint

    Name      string
    Address   string

    CreatedAt time.Time
}