
package config

import (
    "fmt"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

var MasterDB *gorm.DB

func ConnectMasterDB() {
    dsn := "saas_user:password@tcp(127.0.0.1:3306)/master_db?charset=utf8mb4&parseTime=True&loc=Local"
    db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    MasterDB = db
}

func ConnectTenantDB(dsn string) (*gorm.DB, error) {
    return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
