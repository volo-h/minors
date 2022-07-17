package database

import (
  "github.com/volo-h/minors/models"

  "gorm.io/gorm"
  "gorm.io/driver/mysql"
)

var DB *gorm.DB

func Connect() {
  // TODO:
  dsn := "user:password@tcp(127.0.0.1:3306)/db?charset=utf8mb4&parseTime=True&loc=Local"
  database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

  if err != nil {
    panic("Could not connect to the database")
  }

  DB = database

  database.AutoMigrate(&models.User{}, &models.Role{}, &models.Permission{}, &models.Product{}, &models.Order{}, &models.OrderItem{})
  // fmt.Println(db)
}
