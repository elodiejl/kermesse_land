package seeders

import (
	"back/models"
	"gorm.io/gorm"
)

func SeedStands(db *gorm.DB) error {
	stands := []models.Stand{
		{
			KermesseID:        1, // ID de la kermesse correspondante
			Name:              "Stand de Nourriture",
			StandType:         "nourriture",
			ParticipationCost: 3, // Coût en jetons
			TeneurID:          8, // Référence à un utilisateur avec le rôle 'stand leader'
			Stock:             10,
		},
		{
			KermesseID:        1,
			Name:              "Stand de Boisson",
			StandType:         "boisson",
			ParticipationCost: 2,
			TeneurID:          9,
			Stock:             20,
		},
		{
			KermesseID:        2,
			Name:              "Stand d'Activité",
			StandType:         "activité",
			ParticipationCost: 5,
			TeneurID:          10,
			Stock:             0, // Pas de stock pour les activités
		},
	}

	for _, stand := range stands {
		if err := db.Create(&stand).Error; err != nil {
			return err
		}
	}
	return nil
}
