package repositories

import (
	"back/models"
	"gorm.io/gorm"
)

type StandRepository interface {
	Create(stand *models.Stand) error
	FindByID(id uint) (*models.Stand, error)
	FindAllByKermesseID(kermesseID int) ([]models.Stand, error)
	Delete(id uint) error
}

type StandRepositoryImpl struct {
	db *gorm.DB
}

func NewStandRepository(db *gorm.DB) StandRepository {
	return &StandRepositoryImpl{db: db}
}

// Créer un nouveau stand
func (r *StandRepositoryImpl) Create(stand *models.Stand) error {
	return r.db.Create(stand).Error
}

// Trouver un stand par ID
func (r *StandRepositoryImpl) FindByID(id uint) (*models.Stand, error) {
	var stand models.Stand
	if err := r.db.First(&stand, id).Error; err != nil {
		return nil, err
	}
	return &stand, nil
}

// Récupérer tous les stands pour une kermesse
func (r *StandRepositoryImpl) FindAllByKermesseID(kermesseID int) ([]models.Stand, error) {
	var stands []models.Stand
	if err := r.db.Where("kermesse_id = ?", kermesseID).Find(&stands).Error; err != nil {
		return nil, err
	}
	return stands, nil
}

// Supprimer un stand
func (r *StandRepositoryImpl) Delete(id uint) error {
	return r.db.Delete(&models.Stand{}, id).Error
}
