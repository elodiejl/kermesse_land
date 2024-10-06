package repositories

import (
	"back/models"
	"gorm.io/gorm"
)

type PrizeRepository interface {
	Create(prize *models.Prize) error
	FindByID(id uint) (*models.Prize, error)
	FindAllByTombolaID(tombolaID int) ([]models.Prize, error)
	Delete(id uint) error
}

type PrizeRepositoryImpl struct {
	db *gorm.DB
}

func NewPrizeRepository(db *gorm.DB) PrizeRepository {
	return &PrizeRepositoryImpl{db: db}
}

// Créer un nouveau lot
func (r *PrizeRepositoryImpl) Create(prize *models.Prize) error {
	return r.db.Create(prize).Error
}

// Trouver un lot par ID
func (r *PrizeRepositoryImpl) FindByID(id uint) (*models.Prize, error) {
	var prize models.Prize
	if err := r.db.First(&prize, id).Error; err != nil {
		return nil, err
	}
	return &prize, nil
}

// Récupérer tous les lots d'une tombola
func (r *PrizeRepositoryImpl) FindAllByTombolaID(tombolaID int) ([]models.Prize, error) {
	var prizes []models.Prize
	if err := r.db.Where("tombola_id = ?", tombolaID).Find(&prizes).Error; err != nil {
		return nil, err
	}
	return prizes, nil
}

// Supprimer un prix
func (r *PrizeRepositoryImpl) Delete(id uint) error {
	return r.db.Delete(&models.Prize{}, id).Error
}
