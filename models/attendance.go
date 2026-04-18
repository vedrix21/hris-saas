package models

import "time"

type Company struct {
    ID        uint

    Username      string
    CheckIn   time.Time

    Created_at time.Time
}