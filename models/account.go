package models

type Account struct {
    ID           uint
    Code         string `gorm:"uniqueIndex"`

    // 🔥 subscription
    Package      string // basic, pro, premium
    MonthlyFee   int
    UserLimit    int
    IsActive     bool

    // 🔥 billing
    ExpiredAt    *time.Time
    IsTrial      bool

    // 🔥 branding
    LogoURL      string
    ThemeColor   string

    // 🔥 environment
    Environment  string // dev / prod

    // 🔥 feature flags (AI, payroll, dll)
    Features     string // JSON

    CreatedAt    time.Time
    UpdatedAt    time.Time
}