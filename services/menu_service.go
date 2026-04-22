package services

type MenuItem struct {
	Name string
	Link string
}

func GetSidebarByRole(role string) []MenuItem {

	switch role {

	case "owner":
		return []MenuItem{
			{"Dashboard", "/owner/dashboard"},
			{"Clients", "/owner/clients"},
			{"Sales", "/owner/sales"},
			{"Settings", "/owner/settings"},
		}

	case "admin":
		return []MenuItem{
			{"Dashboard", "/dashboard"},
			{"Employees", "/employees"},
			{"Attendance", "/attendance"},
			{"Payroll", "/payroll"},
			{"Settings", "/settings"},
		}

	case "user":
		return []MenuItem{
			{"Dashboard", "/dashboard"},
			{"My Attendance", "/attendance"},
			{"Payslip", "/payslip"},
		}
	}

	return []MenuItem{}
}