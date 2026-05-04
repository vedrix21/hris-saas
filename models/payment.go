package models
import "time"

type Payment struct {
    ID        uint
    AccountID uint
	Account   Account `gorm:"foreignKey:AccountID"`

    Amount    int
    Proof     string // path gambar
    Status    string // pending, approved, rejected
	Note string

    CreatedAt time.Time
}