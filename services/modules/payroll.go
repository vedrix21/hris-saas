package modules

import (
    "fmt"
)

func ProcessPayroll(accountCode string, env string) error {

    fmt.Println("RUN PAYROLL:", accountCode, env)

    // 🔥 simulasi error di dev
    if env == "dev" {
        return fmt.Errorf("Payroll error in dev mode")
    }

    return nil
}