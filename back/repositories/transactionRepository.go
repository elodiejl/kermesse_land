package repositories

import (
	"back/models"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	Create(transaction *models.Transaction) error
	FindByID(id uint) (*models.Transaction, error)
	Update(transaction *models.Transaction) error
	Delete(id uint) error
	FindByParentID(parentID int) ([]models.Transaction, error) // Méthode pour trouver les transactions par ID de parent
}

type TransactionRepositoryImpl struct {
	db *gorm.DB
}

// NewTransactionRepository Retourne une nouvelle instance de TransactionRepository
func NewTransactionRepository(db *gorm.DB) *TransactionRepositoryImpl {
	return &TransactionRepositoryImpl{db: db}
}

// Create Créer une nouvelle transaction
func (r *TransactionRepositoryImpl) Create(transaction *models.Transaction) error {
	return r.db.Create(transaction).Error
}

// FindByID Trouver une transaction par ID
func (r *TransactionRepositoryImpl) FindByID(id uint) (*models.Transaction, error) {
	var transaction models.Transaction
	if err := r.db.First(&transaction, id).Error; err != nil {
		return nil, err
	}
	return &transaction, nil
}

// Update Met à jour les informations d'une transaction
func (r *TransactionRepositoryImpl) Update(transaction *models.Transaction) error {
	return r.db.Save(transaction).Error
}

// Delete Supprimer une transaction par ID
func (r *TransactionRepositoryImpl) Delete(id uint) error {
	return r.db.Delete(&models.Transaction{}, id).Error
}

// FindByParentID Trouver toutes les transactions d'un parent par ID
func (r *TransactionRepositoryImpl) FindByParentID(parentID int) ([]models.Transaction, error) {
	var transactions []models.Transaction
	if err := r.db.Where("parent_id = ?", parentID).Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}
