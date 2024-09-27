package seeders

import (
	"back/models"
	"gorm.io/gorm"
	"time"
)

func SeedKermesses(db *gorm.DB) error {
	kermesses := []models.Kermesse{
		{
			Name:     "Kermesse de l'École A",
			Location: "École A, Paris",
			Date:     time.Now().AddDate(0, 1, 0).Format("2024-09-02"),
		},
		{
			Name:     "Kermesse de l'École B",
			Location: "École B, Lyon",
			Date:     time.Now().AddDate(0, 2, 0).Format("2024-09-10"),
		},
	}

	for _, kermesse := range kermesses {
		if err := db.Create(&kermesse).Error; err != nil {
			return err
		}
	}
	return nil
}
