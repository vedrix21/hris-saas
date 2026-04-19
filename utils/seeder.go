package utils

import (
	"fmt"
	"hris/config"
	"hris/models"

	"golang.org/x/crypto/bcrypt"
)

func SeedOwner() {
	db := config.DB

	var count int64
	db.Model(&models.User{}).
		Where("role = ?", "owner").
		Count(&count)

	// ✅ kalau sudah ada → skip
	if count > 0 {
		fmt.Println("✅ Owner sudah ada, skip seeding")
		return
	}

	fmt.Println("🌱 Creating owner account...")

	// 🔥 generate password random
	rawPassword := GeneratePassword(10)

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(rawPassword), bcrypt.DefaultCost)

	// 🔥 create account dulu (optional kalau belum ada)
	account := models.Account{
		Code:        "aitherhr",
		CompanyName: "Aither HR",
		IsActive:    true,
	}

	db.Create(&account)

	// 🔥 create owner
	user := models.User{
		Username:  "fauzan.control",
		Password:  string(hashedPassword),
		Role:      "owner",
		AccountID: account.ID,
	}

	db.Create(&user)

	fmt.Println("🔥 OWNER CREATED")
	fmt.Println("Username :", user.Username)
	fmt.Println("Password :", rawPassword)

}