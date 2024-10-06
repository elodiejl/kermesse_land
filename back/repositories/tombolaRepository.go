package repositories

import (
	"back/models"
	"gorm.io/gorm"
)

type TombolaRepository interface {
	Create(tombola *models.Tombola) error
	FindByID(id uint) (*models.Tombola, error)
	FindAllByKermesseID(kermesseID int) ([]models.Tombola, error)
	Update(tombola *models.Tombola) error
	Delete(id uint) error
}

type TombolaRepositoryImpl struct {
	db *gorm.DB
}

// NewTombolaRepository crée une nouvelle instance de TombolaRepositoryImpl
func NewTombolaRepository(db *gorm.DB) *TombolaRepositoryImpl {
	return &TombolaRepositoryImpl{db: db}
}

// Create Créer une nouvelle tombola
func (r *TombolaRepositoryImpl) Create(tombola *models.Tombola) error {
	return r.db.Create(tombola).Error
}

// FindByID Trouver une tombola par ID
func (r *TombolaRepositoryImpl) FindByID(id uint) (*models.Tombola, error) {
	var tombola models.Tombola
	if err := r.db.First(&tombola, id).Error; err != nil {
		return nil, err
	}
	return &tombola, nil
}

// FindAllByKermesseID Récupérer toutes les tombolas pour une kermesse spécifique
func (r *TombolaRepositoryImpl) FindAllByKermesseID(kermesseID int) ([]models.Tombola, error) {
	var tombolas []models.Tombola
	if err := r.db.Where("kermesse_id = ?", kermesseID).Find(&tombolas).Error; err != nil {
		return nil, err
	}
	return tombolas, nil
}

// Update Mettre à jour une tombola existante
func (r *TombolaRepositoryImpl) Update(tombola *models.Tombola) error {
	return r.db.Save(tombola).Error
}

// Delete Supprimer une tombola par ID
func (r *TombolaRepositoryImpl) Delete(id uint) error {
	return r.db.Delete(&models.Tombola{}, id).Error
}
