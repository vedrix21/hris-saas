package models

import "time"

type EmployeeSalaryHistory struct {
	ID             uint
	EmployeeID     uint

	BasicSalary    float64

	EffectiveDate  time.Time
	EndDate        *time.Time

	CreatedAt      time.Time
}
