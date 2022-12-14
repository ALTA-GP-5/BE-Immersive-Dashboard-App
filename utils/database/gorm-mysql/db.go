package utils

import (
	"fmt"
	"immersive/config"
	"log"

	classModel "immersive/domains/class/models"
	feedbackModel "immersive/domains/feedback/models"
	menteeModel "immersive/domains/mentee/models"
	mentorModel "immersive/domains/mentor/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(cfg *config.AppConfig) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.DB_USER, cfg.DB_PASS, cfg.DB_HOST, cfg.DB_PORT, cfg.DB_NAME)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	autoMigrate(db)

	return db
}

func autoMigrate(db *gorm.DB) {
	db.AutoMigrate(
		new(mentorModel.Mentor),
		new(classModel.Class),
		new(menteeModel.Mentee),
		new(menteeModel.MenteeDetail),
		new(feedbackModel.FeedBack),
	)
}
