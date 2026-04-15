package services

import (
    "errors"
    "strings"
    "hris/config"
    "hris/models"

    "golang.org/x/crypto/bcrypt"
)



func Login(accountCode, username, password string) (*models.User, *models.Account, error) {
    db := config.DB

    var account models.Account
    if err := db.Where("code = ?", accountCode).First(&account).Error; err != nil {
        return nil, nil, err
    }

    var user models.User
    if err := db.Where("username = ? AND account_id = ?", strings.ToLower(username), account.ID).First(&user).Error; err != nil {
        return nil, nil, err
    }

    // 🔥 GANTI VALIDASI DI SINI (bcrypt)
    err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
    if err != nil {
        return nil, nil, errors.New("invalid password")
    }

    // fmt.Println("🔥 MASUK KE LOGIN SERVICE")

    // fmt.Println("INPUT:")
    // fmt.Println("accountCode:", accountCode)
    // fmt.Println("username:", username)
    // fmt.Println("password:", password)

    // fmt.Println("ACCOUNT FOUND:", account.ID)

    // fmt.Println("DB USER:", user.Username)
    // fmt.Println("DB PASSWORD:", user.Password)

    return &user, &account, nil
}