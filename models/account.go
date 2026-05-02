package models
import "time"

type Account struct {
    ID           uint
    Code string `gorm:"type:varchar(20);uniqueIndex:idx_accounts_code"`

    CompanyName string

    // 🔥 subscription
    Package      string // basic, pro, premium
    MonthlyFee   int
    UserLimit    int
    IsActive     bool
    IsOwner      bool
    SubscriptionStart   *time.Time
    SubscriptionEnd     *time.Time
    LastReminderSent    *time.Time

    // 🔥 billing
    ExpiredAt    *time.Time
    IsTrial      bool

    // 🔥 branding
    LogoURL      string
    ThemeColor   string

    // 🔥 environment
    Environment  string // dev / prod

    // 🔥 feature flags (AI, payroll, dll)
    Features     string `gorm:"type:json"` // JSON

    CreatedAt    time.Time
    UpdatedAt    time.Time
}