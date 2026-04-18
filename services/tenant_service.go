package services

import (
    "hris/config"
    "hris/models"
)

func GetAccountByCode(code string) (*models.Account, error) {
    var account models.Account

    err := config.DB.Where("code = ?", code).First(&account).Error
    if err != nil {
        return nil, err
    }

    return &account, nil
}