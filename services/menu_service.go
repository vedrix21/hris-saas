package services

type MenuItem struct {
	Name     string
	Link     string
	Children []MenuItem
}

func GetSidebarByRole(role string) []MenuItem {

	switch role {

	case "owner":
		return []MenuItem{
			{
				Name: "Dashboard",
				Link: "/owner/dashboard",
			},
			{
				Name: "Clients",
				Children: []MenuItem{
					{"All Clients", "/owner/clients", nil},
					{"Create Client", "/owner/create-account", nil},
					{"Licenses", "/owner/clients/licenses", nil},
					{"Marketing", "/owner/clients/marketing", nil},
				},
			},
			{
				Name: "Sales",
				Link: "/owner/sales",
			},
			{
				Name: "Settings",
				Link: "/owner/settings",
			},
		}

	case "admin":
		return []MenuItem{
			{"Dashboard", "/dashboard", nil},
			{"Employees", "/employees", nil},
			{"Attendance", "/attendance", nil},
			{"Payroll", "/payroll", nil},
			{"Settings", "/settings", nil},
		}

	case "user":
		return []MenuItem{
			{"Dashboard", "/dashboard", nil},
			{"My Attendance", "/attendance", nil},
			{"Payslip", "/payslip", nil},
		}
	}

	return []MenuItem{}
}
