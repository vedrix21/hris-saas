package models
import "time"

type Subscriptionplan struct {
    ID        uint
    
    PlanName  string
    Limituser int
    Price     int

    CreatedAt time.Time
}