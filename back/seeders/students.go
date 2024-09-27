package seeders

import (
	"back/models"
	"gorm.io/gorm"
)

func SeedStudent(db *gorm.DB) error {
	students := []models.Student{
		{
			UserID:      3,
			ParentID:    1,
			TokenAmount: 3,
		},
		{
			UserID:      4,
			ParentID:    1,
			TokenAmount: 2,
		},
		{
			UserID:      5,
			ParentID:    2,
			TokenAmount: 5,
		},
	}

	for _, student := range students {
		if err := db.Create(&student).Error; err != nil {
			return err
		}
	}
	return nil
}
