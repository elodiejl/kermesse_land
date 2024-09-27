package repositories

import (
	"back/models"
	"gorm.io/gorm"
)

type TombolaRepository struct {
	db *gorm.DB
}

func NewTombolaRepository(db *gorm.DB) *TombolaRepository {
	return &TombolaRepository{db: db}
}

// Créer une nouvelle tombola
func (r *TombolaRepository) Create(tombola *models.Tombola) error {
	return r.db.Create(tombola).Error
}

// Trouver une tombola par ID
func (r *TombolaRepository) FindByID(id uint) (*models.Tombola, error) {
	var tombola models.Tombola
	if err := r.db.First(&tombola, id).Error; err != nil {
		return nil, err
	}
	return &tombola, nil
}

// Récupérer toutes les tombolas pour une kermesse
func (r *TombolaRepository) FindAllByKermesseID(kermesseID int) ([]models.Tombola, error) {
	var tombolas []models.Tombola
	if err := r.db.Where("kermesse_id = ?", kermesseID).Find(&tombolas).Error; err != nil {
		return nil, err
	}
	return tombolas, nil
}

// Supprimer une tombola par ID
func (r *TombolaRepository) Delete(id uint) error {
	return r.db.Delete(&models.Tombola{}, id).Error
}
