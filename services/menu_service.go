package services

type MenuItem struct {
	Name string
	Link string
	Children []MenuItem
}

func GetSidebar(role string, tenant string) []MenuItem {

	// 👑 OWNER (NO GATING)
	if role == "owner" {
		return []MenuItem{
			{
				Name: "Dashboard",
				Link: "/owner/dashboard",
			},
			{
				Name: "Create Client",
				Link: "/owner/create_account",
			},
			{
				Name: "List Payments",
				Link: "/owner/payments",
			},
			{
				Name: "Settings",
				Children: []MenuItem{
					{
						Name: "Organization Settings",
						Link: "/owner/settings/organization",
					},
					{
						Name: "System Settings",
						Link: "/owner/settings/system",
					},
				},
			},
		}
	}

	// 🔥 ambil feature tenant
	features, _ := GetFeatures(tenant)

	menu := []MenuItem{
		{
			Name: "Dashboard",
			Link: "/dashboard",
		},
	}

	// 🔥 FEATURE BASED MENU
	if features["employee"] {
		menu = append(menu, MenuItem{
			Name: "Employees",
			Link: "/employees",
		})
	}

	if features["attendance"] {
		menu = append(menu, MenuItem{
			Name: "Attendance",
			Link: "/attendance",
		})
	}

	if features["payroll"] {
		menu = append(menu, MenuItem{
			Name: "Payroll",
			Link: "/payroll",
		})
	}

	// ⚙ SETTINGS MENU
	settingsChildren := []MenuItem{
		{
			Name: "Organization Settings",
			Link: "/settings/organization",
		},
		{
			Name: "Company Settings",
			Link: "/settings/company",
		},
		{
			Name: "Attendance Settings",
			Link: "/settings/attendance",
		},
		{
			Name: "Payroll Settings",
			Link: "/settings/payroll",
		},
	}

	// 🔥 hanya superadmin
	if role == "superadmin" {
		settingsChildren = append(settingsChildren, MenuItem{
			Name: "Data Migration",
			Link: "/settings/data-migration",
		})
	}

	menu = append(menu, MenuItem{
		Name:     "Settings",
		Children: settingsChildren,
	})

	return menu
}
