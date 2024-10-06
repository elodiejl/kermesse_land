package repositories

import (
	"back/models"
	"gorm.io/gorm"
)

type MessageRepository interface {
	SaveMessage(message *models.Message) error
	FindMessagesBySenderAndReceiver(senderID, receiverID int) ([]models.Message, error)
	DeleteMessage(id int) error
}

type MessageRepositoryImpl struct {
	DB *gorm.DB
}

// NewMessageRepository crée une nouvelle instance de MessageRepositoryImpl
func NewMessageRepository(db *gorm.DB) MessageRepository {
	return &MessageRepositoryImpl{DB: db}
}

// SaveMessage Crée un nouveau message de chat
func (repo *MessageRepositoryImpl) SaveMessage(message *models.Message) error {
	return repo.DB.Create(message).Error
}

// FindMessagesBySenderAndReceiver Récupérer tous les messages entre un expéditeur et un récepteur
func (repo *MessageRepositoryImpl) FindMessagesBySenderAndReceiver(senderID, receiverID int) ([]models.Message, error) {
	var messages []models.Message
	if err := repo.DB.Where("sender_id = ? AND receiver_id = ?", senderID, receiverID).Find(&messages).Error; err != nil {
		return nil, err
	}
	return messages, nil
}

// DeleteMessage Supprime un message par ID
func (repo *MessageRepositoryImpl) DeleteMessage(id int) error {
	return repo.DB.Delete(&models.Message{}, id).Error
}
