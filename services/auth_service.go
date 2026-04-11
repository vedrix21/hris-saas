package services

import (
    "hris/config"
    "hris/models"
)

func Login(accountCode, username, password string) (*models.User, *models.Account, error) {
    db := config.DB

    var account models.Account
    if err := db.Where("code = ?", accountCode).First(&account).Error; err != nil {
        return nil, nil, err
    }

    var user models.User
    if err := db.Where("username = ? AND account_id = ?", username, account.ID).First(&user).Error; err != nil {
        return nil, nil, err
    }

    // sementara tanpa hash (nanti kita upgrade)
    if user.Password != password {
        return nil, nil, err
    }

    return &user, &account, nil
}