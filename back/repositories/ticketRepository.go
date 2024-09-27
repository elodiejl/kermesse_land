package repositories

import (
	"back/models"
	"gorm.io/gorm"
)

type TicketRepository struct {
	db *gorm.DB
}

func NewTicketRepository(db *gorm.DB) *TicketRepository {
	return &TicketRepository{db: db}
}

// Create Créer un nouveau ticket
func (r *TicketRepository) Create(ticket *models.Ticket) error {
	return r.db.Create(ticket).Error
}

// FindByID Trouver un ticket par son ID
func (r *TicketRepository) FindByID(id uint) (*models.Ticket, error) {
	var ticket models.Ticket
	if err := r.db.First(&ticket, id).Error; err != nil {
		return nil, err
	}
	return &ticket, nil
}

// FindAllByTombolaID Récupérer tous les tickets pour une tombola spécifique
func (r *TicketRepository) FindAllByTombolaID(tombolaID int) ([]models.Ticket, error) {
	var tickets []models.Ticket
	if err := r.db.Where("tombola_id = ?", tombolaID).Find(&tickets).Error; err != nil {
		return nil, err
	}
	return tickets, nil
}

// Supprimer un ticket
func (r *TicketRepository) Delete(id uint) error {
	return r.db.Delete(&models.Ticket{}, id).Error
}
