package repositories

import (
	"back/models"
	"gorm.io/gorm"
)

type KermesseRepository struct {
	db *gorm.DB
}

func NewKermesseRepository(db *gorm.DB) *KermesseRepository {
	return &KermesseRepository{db: db}
}

// Create Créer une nouvelle kermesse
func (r *KermesseRepository) Create(kermesse *models.Kermesse) error {
	return r.db.Create(kermesse).Error
}

// FindByID Trouver une kermesse par ID
func (r *KermesseRepository) FindByID(id uint) (*models.Kermesse, error) {
	var kermesse models.Kermesse
	if err := r.db.First(&kermesse, id).Error; err != nil {
		return nil, err
	}
	return &kermesse, nil
}

// FindAll Récupérer toutes les kermesses
func (r *KermesseRepository) FindAll() ([]models.Kermesse, error) {
	var kermesses []models.Kermesse
	if err := r.db.Find(&kermesses).Error; err != nil {
		return nil, err
	}
	return kermesses, nil
}

// Delete Supprimer une kermesse par ID
func (r *KermesseRepository) Delete(id int) error {
	return r.db.Delete(&models.Kermesse{}, id).Error
}
