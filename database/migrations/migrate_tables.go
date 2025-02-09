package migrations

import (
	"inkforge/database/models"
	"inkforge/setup"
)

func MigrateTables() {
	setup.DB.AutoMigrate(&models.User{})
}
