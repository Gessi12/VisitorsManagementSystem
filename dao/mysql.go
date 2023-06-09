package dao
import (
	"VisitorsManagementSystem/models"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/callersystem?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err !=  nil {
		fmt.Println(err)
	}

	db.AutoMigrate(&models.Visitor{},models.Admin{})
	return db
}