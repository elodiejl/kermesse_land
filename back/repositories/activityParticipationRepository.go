package repositories

import (
	"back/models"
	"gorm.io/gorm"
)

type ActivityParticipationRepositoryImpl struct {
	DB *gorm.DB
}

// NewActivityParticipationRepositoryImpl crée une nouvelle instance du repository
func NewActivityParticipationRepositoryImpl(db *gorm.DB) *ActivityParticipationRepositoryImpl {
	return &ActivityParticipationRepositoryImpl{DB: db}
}

// Créer une nouvelle participation à une activité
func (repo *ActivityParticipationRepositoryImpl) CreateParticipation(participation *models.ActivityParticipation) error {
	return repo.DB.Create(participation).Error
}

// Trouver une participation par ID
func (repo *ActivityParticipationRepositoryImpl) FindByID(id int) (*models.ActivityParticipation, error) {
	var participation models.ActivityParticipation
	if err := repo.DB.First(&participation, id).Error; err != nil {
		return nil, err
	}
	return &participation, nil
}

// Trouver toutes les participations d'un utilisateur à une activité
func (repo *ActivityParticipationRepositoryImpl) FindAllByUserID(userID int) ([]models.ActivityParticipation, error) {
	var participations []models.ActivityParticipation
	if err := repo.DB.Where("user_id = ?", userID).Find(&participations).Error; err != nil {
		return nil, err
	}
	return participations, nil
}

// Supprimer une participation
func (repo *ActivityParticipationRepositoryImpl) DeleteParticipation(id int) error {
	return repo.DB.Delete(&models.ActivityParticipation{}, id).Error
}
