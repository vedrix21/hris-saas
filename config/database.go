package config

import (
    "fmt"
    "os"

    "hris/models"

    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

var DB *gorm.DB

func ConnectMasterDB() {
    host := os.Getenv("MYSQLHOST")
    port := os.Getenv("MYSQLPORT")
    user := os.Getenv("MYSQLUSER")
    password := os.Getenv("MYSQLPASSWORD")
    dbname := os.Getenv("MYSQLDATABASE")

    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        user, password, host, port, dbname,
    )

    database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("failed to connect database: " + err.Error())
    }

    DB = database

    err = DB.AutoMigrate(&models.Account{}, &models.User{},&models.Company{},&models.Subscription{},)
    if err != nil {
        panic("❌ migration failed: " + err.Error())
    }

    fmt.Println("✅ Database connected & migrated")

}