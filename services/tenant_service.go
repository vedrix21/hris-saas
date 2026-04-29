package services

import (
    "fmt"
    "hris/config"
    "hris/models"
	"hris/utils"    
    "hris/services"

	"golang.org/x/crypto/bcrypt"
)

func GetAccountByCode(code string) (*models.Account, error) {
    var account models.Account

    err := config.DB.Where("code = ?", code).First(&account).Error
    if err != nil {
        return nil, err
    }

    return &account, nil
}

// 🔥 CORE FUNCTION
func CreateTenant(companyName string) (*models.Account, string, string, error) {
    db := config.DB

    var code string

    // 🔥 LOOP SAMPAI DAPET CODE YANG UNIK
    for {
        code = utils.GenerateAccountCode(companyName)

        var count int64
        db.Model(&models.Account{}).
            Where("code = ?", code).
            Count(&count)

        if count == 0 {
            break
        }
    }

    // 🔥 generate password
    rawPassword := utils.GeneratePassword(10)

    hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(rawPassword), bcrypt.DefaultCost)

    // 🔥 create account
    account := models.Account{
        Code:        code,
        CompanyName: companyName,
        IsActive:    true,
        IsOwner:     false,
    }

    if err := db.Create(&account).Error; err != nil {
        return nil, "", "", err
    }

    // 🔥 create super admin user
    user := models.User{
        Username:  code,
        Password:  string(hashedPassword),
        Role:      "superadmin",
        AccountID: account.ID,
    }

    if err := db.Create(&user).Error; err != nil {
        return nil, "", "", err
    }

    // 🔥 SEND EMAIL
    go func() {
        subject := "New Client Account - AitherHR"

        body := fmt.Sprintf(`
        <h3>Client Account Created</h3>
        <p><b>Company:</b> %s</p>
        <p><b>Account Code:</b> %s</p>
        <p><b>Username:</b> %s</p>
        <p><b>Password:</b> %s</p>
        <br>
        <p>Login: https://app.aitherhr.com</p>
        `, companyName, code, code, rawPassword)

        services.SendEmailHTML("fauzanakbarpr@gmail.com", subject, body)
    }()

    return &account, code, rawPassword, nil
}