package repository

import (
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
	"TodoApp/models"

)

func Create (db *gorm.DB, model Todo) error {
	db.Create(&model)
}

func Read (db *gorm.DB, model Todo, id int) error {
	result := db.First(&user, id)

}