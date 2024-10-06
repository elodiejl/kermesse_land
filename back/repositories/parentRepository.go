package repositories

import (
	"back/models"
	"gorm.io/gorm"
)

type ParentRepository interface {
	Create(parent *models.Parent) error
	FindByID(id uint) (*models.Parent, error)
	Update(parent *models.Parent) error
	Delete(id uint) error
}

type ParentRepositoryImpl struct {
	db *gorm.DB
}

// NewParentRepository Retourne une nouvelle instance de ParentRepository
func NewParentRepository(db *gorm.DB) ParentRepository {
	return &ParentRepositoryImpl{db: db}
}

// Create Créer un nouveau parent
func (r *ParentRepositoryImpl) Create(parent *models.Parent) error {
	return r.db.Create(parent).Error
}

// FindByID Trouver un parent par ID
func (r *ParentRepositoryImpl) FindByID(id uint) (*models.Parent, error) {
	var parent models.Parent
	if err := r.db.First(&parent, id).Error; err != nil {
		return nil, err
	}
	return &parent, nil
}

// Update Met à jour les informations d'un parent
func (r *ParentRepositoryImpl) Update(parent *models.Parent) error {
	return r.db.Save(parent).Error
}

// Delete Supprimer un parent par ID
func (r *ParentRepositoryImpl) Delete(id uint) error {
	return r.db.Delete(&models.Parent{}, id).Error
}
