package backend

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitDB initializes the SQLite database
func InitDB(dbPath string) {
	var err error
	DB, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// Migrate the schema
	// DB.AutoMigrate(&model.User{}, &model.Product{})
	log.Println("Database initialized")
}
