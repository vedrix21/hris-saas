package models
import "time"

type Subscription struct {
    ID        uint
    AccountID uint

    FromPlan  string
    ToPlan    string
    Price     int

    CreatedAt time.Time
}