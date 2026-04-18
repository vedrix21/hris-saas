package models

import "time"

type Payroll struct {
    ID          uint
    EmployeeID  uint
    CompanyID   uint

    BasicSalary float64
    Allowance   float64
    Bonus       float64
    Deduction   float64

    TotalSalary float64
    Status      string // pending, paid

    Period      string // 2026-04

    CreatedAt   time.Time
    UpdatedAt   time.Time
}