package model

import (
    "gorm.io/gorm"
    "gorm.io/driver/postgres"
)

var DB *gorm.DB

func ConnectDB() {
    dsn := "host=localhost user=postgres password=re0800za dbname=golang-final-task port=5432 sslmode=disable TimeZone=Asia/Shanghai"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic(err)
    }

    err = db.AutoMigrate(&User{}, &Photo{})
    if err != nil {
        panic(err)
    }

    DB = db
}

