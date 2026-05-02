package models
import "time"

type Subscriptionplan struct {
    ID        uint
    Code string
    
    PlanName  string
    Limituser int
    Price     int
    Description string

    CreatedAt time.Time
}