package repositories

import (
	"back/models"
	"gorm.io/gorm"
)

type TokenRepository interface {
	GetTokenByID(id uint) (*models.Token, error)
	CreateToken(token *models.Token) error
	UpdateToken(token *models.Token) error
	DeleteToken(id uint) error
	GetTokensByStudentID(studentID int) ([]models.Token, error)
	GetTokensByParentID(parentID int) ([]models.Token, error)
}

type TokenRepositoryImpl struct {
	db *gorm.DB
}

// NewTokenRepository crée une nouvelle instance de TokenRepository
func NewTokenRepository(DB *gorm.DB) *TokenRepositoryImpl {
	return &TokenRepositoryImpl{db: DB}
}

func (r *TokenRepositoryImpl) GetTokenByID(id uint) (*models.Token, error) {
	var token models.Token
	err := r.db.First(&token, id).Error
	if err != nil {
		return nil, err
	}
	return &token, nil
}

func (r *TokenRepositoryImpl) CreateToken(token *models.Token) error {
	return r.db.Create(token).Error
}

func (r *TokenRepositoryImpl) UpdateToken(token *models.Token) error {
	return r.db.Save(token).Error
}

func (r *TokenRepositoryImpl) DeleteToken(id uint) error {
	return r.db.Delete(&models.Token{}, id).Error
}

func (r *TokenRepositoryImpl) GetTokensByStudentID(studentID int) ([]models.Token, error) {
	var tokens []models.Token
	err := r.db.Where("student_id = ?", studentID).Find(&tokens).Error
	return tokens, err
}

func (r *TokenRepositoryImpl) GetTokensByParentID(parentID int) ([]models.Token, error) {
	var tokens []models.Token
	err := r.db.Where("parent_id = ?", parentID).Find(&tokens).Error
	return tokens, err
}
