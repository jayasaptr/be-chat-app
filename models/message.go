package models

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	SenderID   uint   `gorm:"not null"`
	ReceiverID uint   `gorm:"not null"`
	Content    string `gorm:"type:text;not null"`
}
