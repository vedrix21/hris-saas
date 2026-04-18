package models

import "time"

type Attendance struct {
    ID        uint

    Username      string
    CheckIn   time.Time

    Created_at time.Time
}