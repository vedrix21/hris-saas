package services

import "hris/models"

// 🔥 ambil account berdasarkan code (tenant)
func GetAccountByCode(code string) (*models.Account, error) {
    var account models.Account
    err := DB.Where("code = ?", code).First(&account).Error
    if err != nil {
        return nil, err
    }
    return &account, nil
}