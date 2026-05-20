package models

import "time"

type Attendance struct {
    ID        uint
    AccountCode string // e.g. "TENANT001"
    Username      string
    CheckIn   time.Time

    Created_at time.Time
}