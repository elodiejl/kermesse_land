package seeders

import (
	"back/models"
	"gorm.io/gorm"
)

func SeedParent(db *gorm.DB) error {
	parents := []models.Parent{
		{
			UserID:       6,
			TokensAmount: 2,
		},
		{
			UserID:       7,
			TokensAmount: 0,
		},
	}

	for _, parent := range parents {
		if err := db.Create(&parent).Error; err != nil {
			return err
		}
	}
	return nil
}
