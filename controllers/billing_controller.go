package controllers

import (
	
	"hris/config"
	"hris/models"
	"hris/services/modules"
	

	"github.com/gin-gonic/gin"
	
)

func UpgradePlan(c *gin.Context) {
    accountCode, _ := c.Cookie("tenant")

    var acc models.Account
    config.DB.Where("code = ?", accountCode).First(&acc)

    newPlan := c.PostForm("plan")

    config.DB.Model(&acc).Update("package", newPlan)

    config.DB.Create(&models.Subscription{
        AccountID: acc.ID,
        FromPlan:  acc.Package,
        ToPlan:    newPlan,
    })

    c.String(200, "Plan updated")
}