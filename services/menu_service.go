package services

type MenuItem struct {
	Name string
	Link string
	// Children []MenuItem
}

func GetSidebar(role string, tenant string) []MenuItem {

	// 👑 OWNER (NO GATING)
	if role == "owner" {
		return []MenuItem{
			{"Dashboard", "/owner/dashboard"},
			{"Create Client", "/owner/create_account"},
			{"List Payments", "/owner/payments"},
			{"Settings", "/owner/settings"},
		}
	}

	// 🔥 ambil features dari DB
	features, _ := GetFeatures(tenant)

	menu := []MenuItem{
		{"Dashboard", "/dashboard"},
	}

	// 🔥 FEATURE BASED MENU
	if features["employee"] {
		menu = append(menu, MenuItem{"Employees", "/employees"})
	}

	if features["attendance"] {
		menu = append(menu, MenuItem{"Attendance", "/attendance"})
	}

	if features["payroll"] {
		menu = append(menu, MenuItem{"Payroll", "/payroll"})
	}

	// nanti future
	// if features["loan"] { ... }
	// if features["reimbursement"] { ... }

	menu = append(menu, MenuItem{"Settings", "/settings"})

	return menu
}
