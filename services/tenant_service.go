package services

import (
	// "fmt"
	"encoding/json"
	"hris/config"
	"hris/models"
	"hris/utils"

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
func CreateTenant(companyName string, plan models.Subscriptionplan, picname string, picemail string) (*models.Account, string, string, error) {
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

	planConfig := GetPlanConfig(plan.Code)

	var subplan models.Subscriptionplan
	db.Model(&models.Subscriptionplan{}).Where("plan_name = ?", planConfig.Name).First(&subplan)

	featuresJSON, _ := json.Marshal(planConfig.Features)

	// 🔥 create account
	account := models.Account{
		Code:        code,
		CompanyName: companyName,
		IsActive:    true,
		IsOwner:     false,
		Package:     planConfig.Name,
		Features:    string(featuresJSON),
		MonthlyFee:  subplan.Price,
		UserLimit:   subplan.Limituser,
		PicName:     picname,
		PicEmail:    picemail,
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

	var companyID string

	for {
		companyID = utils.GenerateCompanyID()

		var count int64
		db.Model(&models.Company{}).
			Where("company_id = ?", companyID).
			Count(&count)

		if count == 0 {
			break
		}
	}

	// 🔥 create first company
	company := models.Company{
		AccountID: account.ID,
		CompanyID: companyID,
		Name:      companyName,
		Address:   "Default Address",
	}

	if err := db.Create(&company).Error; err != nil {
		return nil, "", "", err
	}

	// 🔥 SEND EMAIL
	// go func() {
	//     subject := "New Client Account - AitherHR"

	//     body := fmt.Sprintf(`
	//     <h3>Client Account Created</h3>
	//     <p><b>Company:</b> %s</p>
	//     <p><b>Account Code:</b> %s</p>
	//     <p><b>Username:</b> %s</p>
	//     <p><b>Password:</b> %s</p>
	//     <br>
	//     <p>Login: https://app.aitherhr.com</p>
	//     `, companyName, code, code, rawPassword)

	//     SendEmailHTML("fauzanakbarpr@gmail.com", subject, body)
	// }()

	return &account, code, rawPassword, nil
}
