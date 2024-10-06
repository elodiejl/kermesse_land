package seeders

import (
	"back/models"
	"gorm.io/gorm"
)

func SeedKermesses(db *gorm.DB) error {
	kermesses := []models.Kermesse{
		{
			Name:     "Kermesse de l'École A",
			Location: "École A, Paris",
			Date:     "2024-10-12",
		},
		{
			Name:     "Kermesse de l'École B",
			Location: "École B, Lyon",
			Date:     "2024-10-16",
		},
	}

	for _, kermesse := range kermesses {
		if err := db.Create(&kermesse).Error; err != nil {
			return err
		}
	}
	return nil
}
