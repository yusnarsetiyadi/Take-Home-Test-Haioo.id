package mysql

import (
	"fmt"
	"log"

	"haiooId/shoppingChart/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(cfg *config.AppConfig) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True", cfg.DB_USERNAME, cfg.DB_PASSWORD, cfg.DB_HOST, cfg.DB_PORT, cfg.DB_NAME)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Cannot connect to DB")
	}

	// MigrateDB(db)

	return db
}

// func MigrateDB(db *gorm.DB) {
// 	db.AutoMigrate(&user.User{})
// 	db.AutoMigrate(&homestay.Homestay{})
// 	db.AutoMigrate(&feedback.Feedback{})
// 	db.AutoMigrate(&reservation.Reservation{})
// }
