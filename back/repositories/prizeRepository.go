package repositories

import (
	"back/models"
	"gorm.io/gorm"
)

type PrizeRepository struct {
	db *gorm.DB
}

func NewPrizeRepository(db *gorm.DB) *PrizeRepository {
	return &PrizeRepository{db: db}
}

// Créer un nouveau lot
func (r *PrizeRepository) Create(prize *models.Prize) error {
	return r.db.Create(prize).Error
}

// Trouver un lot par ID
func (r *PrizeRepository) FindByID(id uint) (*models.Prize, error) {
	var prize models.Prize
	if err := r.db.First(&prize, id).Error; err != nil {
		return nil, err
	}
	return &prize, nil
}

// Récupérer tous les lots d'une tombola
func (r *PrizeRepository) FindAllByTombolaID(tombolaID int) ([]models.Prize, error) {
	var prizes []models.Prize
	if err := r.db.Where("tombola_id = ?", tombolaID).Find(&prizes).Error; err != nil {
		return nil, err
	}
	return prizes, nil
}
