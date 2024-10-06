package repositories

import (
	"back/models"
	"gorm.io/gorm"
)

type TicketRepository interface {
	Create(ticket *models.Ticket) error
	FindByID(id uint) (*models.Ticket, error)
	FindAllByTombolaID(tombolaID int) ([]models.Ticket, error)
	FindAllByStudentID(studentID int) ([]models.Ticket, error) // Nouvelle méthode
	Delete(id uint) error
}

type TicketRepositoryImpl struct {
	db *gorm.DB
}

// NewTicketRepository crée une nouvelle instance de TicketRepositoryImpl
func NewTicketRepository(db *gorm.DB) *TicketRepositoryImpl {
	return &TicketRepositoryImpl{db: db}
}

// Create Créer un nouveau ticket
func (r *TicketRepositoryImpl) Create(ticket *models.Ticket) error {
	return r.db.Create(ticket).Error
}

// FindByID Trouver un ticket par son ID
func (r *TicketRepositoryImpl) FindByID(id uint) (*models.Ticket, error) {
	var ticket models.Ticket
	if err := r.db.First(&ticket, id).Error; err != nil {
		return nil, err
	}
	return &ticket, nil
}

// FindAllByTombolaID Récupérer tous les tickets pour une tombola spécifique
func (r *TicketRepositoryImpl) FindAllByTombolaID(tombolaID int) ([]models.Ticket, error) {
	var tickets []models.Ticket
	if err := r.db.Where("tombola_id = ?", tombolaID).Find(&tickets).Error; err != nil {
		return nil, err
	}
	return tickets, nil
}

// FindAllByStudentID Récupérer tous les tickets par StudentID
func (r *TicketRepositoryImpl) FindAllByStudentID(studentID int) ([]models.Ticket, error) {
	var tickets []models.Ticket
	if err := r.db.Where("student_id = ?", studentID).Find(&tickets).Error; err != nil {
		return nil, err
	}
	return tickets, nil
}

// Delete Supprimer un ticket par son ID
func (r *TicketRepositoryImpl) Delete(id uint) error {
	return r.db.Delete(&models.Ticket{}, id).Error
}
