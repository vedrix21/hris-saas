package services

type Plan struct {
	Name     string
	Features map[string]bool
}

// 🔥 MASTER PLAN CONFIG
func GetPlanConfig(planName string) Plan {

	switch planName {

	case "basic":
		return Plan{
			Name: "Basic",
			Features: map[string]bool{
				"employee":   true,
				"attendance": true,
				"payroll":    false,
				"loan":       false,
				"reimbursement": false,
				"analytics":  false,
				"api":        false,
			},
		}

	case "pro":
		return Plan{
			Name: "Pro",
			Features: map[string]bool{
				"employee":   true,
				"attendance": true,
				"payroll":    true,
				"loan":       true,
				"reimbursement": true,
				"analytics":  false,
				"api":        false,
			},
		}

	case "enterprise":
		return Plan{
			Name: "Enterprise",
			Features: map[string]bool{
				"employee":   true,
				"attendance": true,
				"payroll":    true,
				"loan":       true,
				"reimbursement": true,
				"analytics":  true,
				"api":        true,
			},
		}
	}

	return Plan{}
}