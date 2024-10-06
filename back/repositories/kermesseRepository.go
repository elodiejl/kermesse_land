package repositories

import (
	"back/models"
	"gorm.io/gorm"
)

type KermesseRepository interface {
	Create(kermesse *models.Kermesse) error
	FindByID(id uint) (*models.Kermesse, error)
	FindAll() ([]models.Kermesse, error)
	Update(kermesse *models.Kermesse) error
	Delete(id uint) error
}

type KermesseRepositoryImpl struct {
	db *gorm.DB
}

// NewKermesseRepository crée une nouvelle instance de KermesseRepositoryImpl
func NewKermesseRepository(db *gorm.DB) KermesseRepository {
	return &KermesseRepositoryImpl{db: db}
}

// Create Créer une nouvelle kermesse
func (r *KermesseRepositoryImpl) Create(kermesse *models.Kermesse) error {
	return r.db.Create(kermesse).Error
}

// FindByID Trouver une kermesse par ID
func (r *KermesseRepositoryImpl) FindByID(id uint) (*models.Kermesse, error) {
	var kermesse models.Kermesse
	if err := r.db.First(&kermesse, id).Error; err != nil {
		return nil, err
	}
	return &kermesse, nil
}

// FindAll Récupérer toutes les kermesses
func (r *KermesseRepositoryImpl) FindAll() ([]models.Kermesse, error) {
	var kermesses []models.Kermesse
	if err := r.db.Find(&kermesses).Error; err != nil {
		return nil, err
	}
	return kermesses, nil
}

// Update Mettre à jour une kermesse existante
func (r *KermesseRepositoryImpl) Update(kermesse *models.Kermesse) error {
	return r.db.Save(kermesse).Error
}

// Delete Supprimer une kermesse par ID
func (r *KermesseRepositoryImpl) Delete(id uint) error {
	return r.db.Delete(&models.Kermesse{}, id).Error
}
