package repositories

import (
	"back/models"
	"gorm.io/gorm"
)

type ActivityRepositoryImpl struct {
	DB *gorm.DB
}

// NewActivityRepositoryImpl crée une nouvelle instance du repository
func NewActivityRepositoryImpl(db *gorm.DB) *ActivityRepositoryImpl {
	return &ActivityRepositoryImpl{DB: db}
}

// Créer une nouvelle activité
func (repo *ActivityRepositoryImpl) CreateActivity(activity *models.Activity) error {
	return repo.DB.Create(activity).Error
}

// Trouver une activité par ID
func (repo *ActivityRepositoryImpl) FindByID(id int) (*models.Activity, error) {
	var activity models.Activity
	if err := repo.DB.First(&activity, id).Error; err != nil {
		return nil, err
	}
	return &activity, nil
}

// Mettre à jour une activité
func (repo *ActivityRepositoryImpl) UpdateActivity(activity *models.Activity) error {
	return repo.DB.Save(activity).Error
}

// Supprimer une activité
func (repo *ActivityRepositoryImpl) DeleteActivity(id int) error {
	return repo.DB.Delete(&models.Activity{}, id).Error
}

// Trouver toutes les activités pour un stand spécifique
func (repo *ActivityRepositoryImpl) FindAllByStandID(standID int) ([]models.Activity, error) {
	var activities []models.Activity
	if err := repo.DB.Where("stand_id = ?", standID).Find(&activities).Error; err != nil {
		return nil, err
	}
	return activities, nil
}
