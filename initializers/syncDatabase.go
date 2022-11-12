package initializers

import "github.com/GabrielEdwinSP/GolangDeveloperTest/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{}, &models.JobList{})
}
