package initializers

import "github.com/hazaloolu/jwt_auth/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
