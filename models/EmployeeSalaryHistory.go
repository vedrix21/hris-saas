package models

import "time"

type EmployeeSalaryHistory struct {
	ID             uint
	EmployeeID     uint
	AccountCode    string // e.g. "TENANT001"
	BasicSalary    float64

	EffectiveDate  time.Time
	EndDate        *time.Time

	CreatedAt      time.Time
}
