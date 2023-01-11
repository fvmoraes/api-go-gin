package initializers

import "api-go-gin/models"

func StartMigration() {
	DB.AutoMigrate(&models.Foobar{})
}
