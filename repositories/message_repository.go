package repositories

import (
	"chat-app/config"
	"chat-app/models"
)

func CreateMessage(message *models.Message) error {
	return config.DB.Create(message).Error
}

func GetMessages(senderID, receiverID uint) ([]models.Message, error) {
	var messages []models.Message
	err := config.DB.Where("(sender_id = ? AND receiver_id = ?) OR (sender_id = ? AND receiver_id = ?)", senderID, receiverID, receiverID, senderID).
		Find(&messages).Error
	return messages, err
}
