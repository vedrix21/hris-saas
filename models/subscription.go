package models
import "time"

type Subscription struct {
    ID        uint
    AccountID uint

    Account   Account `gorm:"foreignKey:AccountID"`

    FromPlan  string
    ToPlan    string
    Price     int

    Status   string // pending, active
	Proof    string // 🔥 path bukti transfer

    CreatedAt time.Time
}