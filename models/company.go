package models

import "time"

type Company struct {
    ID        uint
    AccountID uint

    Name      string
    Address   string

    CreatedAt time.Time
}