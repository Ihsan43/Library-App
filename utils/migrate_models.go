package utils

import (
	"library_app/model"
	"log"

	"gorm.io/gorm"
)

var mainModels = []interface{}{
	&model.User{},
	&model.Book{},
	&model.Address{},
	// &model.Orders{},
}

func MigrateModels(db *gorm.DB) {
	for _, model := range mainModels {
		err := db.AutoMigrate(model)
		if err != nil {
			log.Fatalf("Error migrating model %v: %s", model, err.Error()) 
		} else {
			log.Printf("Successfully migrated model %v", model)
		}

	}
}
