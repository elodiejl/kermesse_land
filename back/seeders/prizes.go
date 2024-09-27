package seeders

import (
	"back/models"
	"gorm.io/gorm"
)

func SeedPrizes(db *gorm.DB) error {
	prizes := []models.Prize{
		{
			TombolaID: 1,
			Name:      "Bicycle",
		},
		{
			TombolaID: 1,
			Name:      "Télévision",
			//WinnerID:  0, // Aucun gagnant encore (pas encore tiré)
		},
		{
			TombolaID: 2,
			Name:      "Console de jeux",
			//WinnerID:  0,
		},
	}

	for _, prize := range prizes {
		if err := db.Create(&prize).Error; err != nil {
			return err
		}
	}
	return nil
}
