package services

import (
	"chat-app/models"
	"chat-app/repositories"
	"chat-app/utils"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(username, email, password string) error {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	user := &models.User{
		Username: username,
		Email:    email,
		Password: string(hashedPassword),
	}

	return repositories.CreateUser(user)
}

func AuthenticateUser(email, password string) (string, error) {
	user, err := repositories.GetUserByEmail(email)
	if err != nil {
		return "", errors.New("user not found")
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return "", errors.New("invalid password")
	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return token, nil
}
