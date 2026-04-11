package config

import (
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

var DB *gorm.DB

func ConnectMasterDB() {
    dsn := "user:password@tcp(host:port)/dbname?charset=utf8mb4&parseTime=True&loc=Local"

    database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    DB = database
}