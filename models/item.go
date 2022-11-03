package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	Id   int    `json:"id" gorm:"primary_key"`
	Tags string `json:"tags"`
	Data string `json:"data"`
	// UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type CreateItem struct {
	Tags string `json:"tags" binding:"required"`
	Data string `json:"data" binding:"required"`
}

type UpdateItem struct {
	Tags string `json:"tags"`
	Data string `json:"data"`
}

var DB *gorm.DB

const DB_FILE = "./goko.db"

func ConnectDatabase() {
	db, err := gorm.Open(sqlite.Open(DB_FILE), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Item{})

	DB = db
}
