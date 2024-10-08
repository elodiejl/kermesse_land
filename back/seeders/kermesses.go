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
			Date:     "2024-10-01",
		},
		{
			Name:     "Kermesse de l'École B",
			Location: "École B, Lyon",
			Date:     "2024-10-02",
		},
		{
			Name:     "Kermesse de l'École C",
			Location: "École C, Paris",
			Date:     "2024-10-10",
		},
		{
			Name:     "Kermesse de l'École D",
			Location: "École D, Lyon",
			Date:     "2024-10-16",
		},
	}

	for _, kermesse := range kermesses {
		var existingKermesse models.Kermesse
		// Cherche la kermesse par son nom
		if err := db.Where("name = ?", kermesse.Name).First(&existingKermesse).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				// Si la kermesse n'existe pas, crée un nouvel enregistrement
				if err := db.Create(&kermesse).Error; err != nil {
					return err
				}
			} else {
				// Si une autre erreur se produit
				return err
			}
		} else {
			// Si la kermesse existe, mets à jour ses informations
			existingKermesse.Location = kermesse.Location
			existingKermesse.Date = kermesse.Date

			if err := db.Save(&existingKermesse).Error; err != nil {
				return err
			}
		}
	}
	return nil
}
