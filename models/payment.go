package models
import "time"

type Payment struct {
    ID        uint
    AccountID uint

    Amount    int
    Proof     string // path gambar

    Status    string // pending, approved, rejected

    CreatedAt time.Time
}