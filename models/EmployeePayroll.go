package models

import "time"

type EmployeePayroll struct {
	ID                    uint
	EmployeeID            uint

	MaritalStatus         string

	BankName              string
	BankAccountNo         string
	BankHolder            string

	NPWP                  string
	BPJSKetenagakerjaan   string
	BPJSKesehatan         string

	CreatedAt             time.Time
	UpdatedAt             time.Time
}
