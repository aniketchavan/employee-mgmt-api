package persistence

import (
	datamodel "example.com/employee-mgmt/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitMySQL(dsn string) error {
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	// Migrate the schema
	return DB.AutoMigrate(&datamodel.Employee{})
}

func StoreEmployee(employee datamodel.Employee) error {
	return DB.Create(&employee).Error
}
