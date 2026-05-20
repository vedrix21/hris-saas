package models

import "time"

type Employee struct {
	ID               uint
	AccountCode      string // e.g. "TENANT001"
	EmployeeID       string // e.g. "EMP001"
	EmployeeCode     string // e.g. "E001"
	FirstName        string
	MiddleName       string
	LastName         string
	FullName         string
	BirthPlace       string
	BirthDate        string
	Gender           string
	Religion         string
	Phone            string
	Email            string
	Address          string
	Position         string
	Department       string
	JoinDate         time.Time
	EmploymentStatus string // e.g. "Contract", "Permanent"
	EmployeeStatus   string // e.g. "Active", "Inactive"
	EmergencyContactName  string
	EmergencyContactPhone string
	MaritalStatus         string
	Nationality           string
	Photo                 string
	CreatedAt             time.Time
	UpdatedAt             time.Time
}
