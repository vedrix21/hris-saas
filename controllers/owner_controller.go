package controllers

import (
	"fmt"
	"hris/config"
	"hris/models"
	"hris/services"
	"hris/utils"

	"github.com/gin-gonic/gin"
)

func ShowCreateAccount(c *gin.Context) {

	tenant, _ := c.Cookie("tenant")
	user := c.MustGet("user").(models.User)
	menus := services.GetSidebar(user.Role, tenant)

	success := c.Query("success")
	accountCode := c.Query("code")
	username := c.Query("user")
	password := c.Query("pass")

	var plans []models.Subscriptionplan
	config.DB.Find(&plans)

	utils.Render(c, []string{
		"templates/owner/create_account.html",
	}, gin.H{
		"title":       "Create Account",
		"Menus":       menus,
		"CurrentPath": c.Request.URL.Path,
		"success":     success,
		"accountCode": accountCode,
		"username":    username,
		"password":    password,
		"plans":       plans,
	})
}

func ShowSettings(c *gin.Context) {
	tenant, _ := c.Cookie("tenant")
	user := c.MustGet("user").(models.User)
	menus := services.GetSidebar(user.Role, tenant)
	utils.Render(c, []string{
		"templates/owner/settings.html",
	}, gin.H{
		"title":       "Owner Settings",
		"Menus":       menus,
		"CurrentPath": c.Request.URL.Path,
	})
}

func CreateAccount(c *gin.Context) {
	companyName := c.PostForm("company_name")
	picname := c.PostForm("picname")
	picemail := c.PostForm("picemail")
	planID := c.PostForm("plan_id")

	var plan models.Subscriptionplan
	if err := config.DB.Where("code = ?", planID).First(&plan).Error; err != nil {
		c.String(400, "Invalid plan")
		return
	}

	account, username, password, err := services.CreateTenant(companyName, plan, picname, picemail)
	if err != nil {
		c.String(500, fmt.Sprintf("ERROR: %v", err))
		return
	}

	// 🔥 Email content
	body := `
	<h2>New Client Created 🚀</h2>
	<p><b>Company:</b> ` + companyName + `</p>
	<p><b>Account Code:</b> ` + account.Code + `</p>
	<p><b>Username:</b> ` + username + `</p>
	<p><b>Password:</b> ` + password + `</p>

	<br>
	<a href="https://app.aitherhr.com">Login AitherHR</a>
	`

	// 🔥 kirim ke email kamu
	err = services.SendEmailHTML(
		"fauzanakbarpr@gmail.com",
		"New Client AitherHR",
		body,
	)
	if err != nil {
		c.String(500, "Account created but email failed")
		return
	}

	// 🔥 REDIRECT
	c.Redirect(302,
		"/owner/create_account?success=1&code="+account.Code+
			"&user="+username+
			"&pass="+password,
	)
}

func OwnerDashboard(c *gin.Context) {
	var accounts []models.Account

	config.DB.Find(&accounts)

	success := c.Query("success")
	db := config.DB

	var totalClients int64
	db.Model(&models.Account{}).
		Where("is_owner = ?", false).
		Count(&totalClients)

	tenant, _ := c.Cookie("tenant")
	user := c.MustGet("user").(models.User)
	menus := services.GetSidebar(user.Role, tenant)

	utils.Render(c, []string{
		"templates/owner/dashboard.html",
	}, gin.H{
		"title":        "Owner Dashboard",
		"Menus":        menus,
		"CurrentPath":  c.Request.URL.Path,
		"accounts":     accounts,
		"success":      success,
		"TotalClients": totalClients,
	})
}

func SaveSettings(c *gin.Context) {
	logo := c.PostForm("logo_url")
	color := c.PostForm("theme_color")

	tenant, _ := c.Cookie("tenant")

	config.DB.Model(&models.Account{}).
		Where("code = ?", tenant).
		Updates(map[string]interface{}{
			"logo_url":    logo,
			"theme_color": color,
		})

	c.Redirect(302, "/owner/dashboard")
}
