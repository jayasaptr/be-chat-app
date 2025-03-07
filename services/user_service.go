package services

import (
	"chat-app/models"
	"chat-app/repositories"
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

func AuthenticateUser(email, password string) (*models.User, error) {
	user, err := repositories.GetUserByEmail(email)
	if err != nil {
		return nil, errors.New("user not found")
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return nil, errors.New("invalid password")
	}

	return user, nil
}
