package db

import (
	"fmt"
	"gaunRumaRestApi/config"
	"gaunRumaRestApi/entity"
	golog "log"
	"os"

	"github.com/labstack/gommon/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Handler struct {
	DB *gorm.DB
}

func Init(conf config.Configuration) Handler {

	gormlog := logger.New(
		golog.New(os.Stdout, "\r\n", golog.LstdFlags),
		logger.Config{
			LogLevel:                  logger.Info,
			Colorful:                  true,
			IgnoreRecordNotFoundError: true,
		},
	)

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		conf.DB_HOST,
		conf.DB_PORT,
		conf.DB_USERNAME,
		conf.DB_PASSWORD,
		conf.DB_NAME,
		"disable")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: gormlog})
	if err != nil {
		log.Fatal("Failed to open database connection: ", err)
	}
	db.AutoMigrate(&entity.User{}, &entity.ProductType{}, &entity.Product{})
	return Handler{db}
}
