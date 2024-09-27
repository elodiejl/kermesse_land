package seeders

import (
	"back/models"
	"back/services"
	"gorm.io/gorm"
	"time"
)

func SeedTickets(db *gorm.DB) error {
	tickets := []models.Ticket{
		{
			StudentID:    1, // Référence à l'ID d'un élève existant
			KermesseID:   1, // Référence à une kermesse existante
			TombolaID:    1, // Référence à une tombola existante
			TicketNumber: services.GenerateTicketNumber(),
			PurchasedAt:  time.Now().Format("2024-09-01 15:04:05"),
		},
		{
			StudentID:    2,
			KermesseID:   1,
			TombolaID:    1,
			TicketNumber: services.GenerateTicketNumber(),
			PurchasedAt:  time.Now().Format("2024-09-01 15:04:05"),
		},
		{
			StudentID:    3,
			KermesseID:   2,
			TombolaID:    2,
			TicketNumber: services.GenerateTicketNumber(),
			PurchasedAt:  time.Now().Format("2024-09-01 15:04:05"),
		},
		{
			StudentID:    4,
			KermesseID:   2,
			TombolaID:    2,
			TicketNumber: services.GenerateTicketNumber(),
			PurchasedAt:  time.Now().Format("2024-09-02 15:04:05"),
		},
	}

	for _, ticket := range tickets {
		if err := db.Create(&ticket).Error; err != nil {
			return err
		}
	}
	return nil
}
