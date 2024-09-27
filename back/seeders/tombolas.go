package seeders

import (
	"back/models"
	"gorm.io/gorm"
)

func SeedTombolas(db *gorm.DB) error {
	tombolas := []models.Tombola{
		{
			KermesseID: 1,            // Référence à une kermesse existante
			DrawnAt:    "2024-09-01", // Date du tirage
			WinnerID:   3,
			TicketId:   1,
		},
		{
			KermesseID: 2,
			DrawnAt:    "2024-09-15",
			WinnerID:   0,
			TicketId:   0,
		},
	}

	for _, tombola := range tombolas {
		if err := db.Create(&tombola).Error; err != nil {
			return err
		}
	}
	return nil
}
