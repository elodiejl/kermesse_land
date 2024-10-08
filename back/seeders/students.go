package seeders

import (
	"back/models"
	"errors"
	"gorm.io/gorm"
)

func SeedStudent(db *gorm.DB) error {
	students := []models.Student{
		{
			UserID:      3,
			ParentID:    6,
			TokenAmount: 3,
		},
		{
			UserID:      4,
			ParentID:    6,
			TokenAmount: 2,
		},
		{
			UserID:      5,
			ParentID:    7,
			TokenAmount: 5,
		},
	}

	for _, student := range students {
		var existingStudent models.Student

		// Rechercher l'étudiant existant par UserID
		if err := db.Where("user_id = ?", student.UserID).First(&existingStudent).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				// Si l'étudiant n'existe pas, on le crée
				if err := db.Create(&student).Error; err != nil {
					return err
				}
			} else {
				return err
			}
		} else {
			existingStudent.ParentID = student.ParentID
			existingStudent.TokenAmount = student.TokenAmount

			if err := db.Save(&existingStudent).Error; err != nil {
				return err
			}
		}
	}
	return nil
}
