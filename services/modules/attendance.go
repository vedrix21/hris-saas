package modules

import "fmt"

func ProcessAttendance(accountCode string, env string) error {
    
    fmt.Println("RUN ATTENDANCE:", accountCode, env)

    // simulasi error di dev
    if env == "dev" {
        return fmt.Errorf("Attendance error in dev mode")
    }

    return nil
}