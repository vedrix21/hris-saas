package controllers

import (
	"fmt"
	"hris/config"
	"hris/models"
	"hris/services"
	"hris/utils"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func UpgradePlan(c *gin.Context) {
    accountCode, _ := c.Cookie("tenant")

    var acc models.Account
    config.DB.Where("code = ?", accountCode).First(&acc)

    newPlan := c.PostForm("plan")

    config.DB.Model(&acc).Update("package", newPlan)

    config.DB.Create(&models.SubscriptionHistory{
        AccountID: acc.ID,
        FromPlan:  acc.Package,
        ToPlan:    newPlan,
    })

    c.String(200, "Plan updated")
}