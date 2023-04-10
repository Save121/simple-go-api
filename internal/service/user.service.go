package service

import (
	"context"
	"errors"

	"github.com/Save121/simple-go-api/encryption"
	"github.com/Save121/simple-go-api/internal/models"
)

var (
	ErrUserAlreadyExists  = errors.New("user already exist")
	ErrInvalidCredentials = errors.New("invalid credentials")
)

func (s *serv) RegisterUser(ctx context.Context, email, name, password string) error {
	if u, _ := s.repo.GetUserByEmail(ctx, email); u != nil {
		return ErrUserAlreadyExists
	}
	pass, err := encryption.Encrypt([]byte(password))
	if err != nil {
		return err
	}
	encryptedPassword := encryption.ToBase64(pass)
	return s.repo.SaveUser(ctx, email, name, encryptedPassword)
}

func (s *serv) LoginUser(ctx context.Context, email, password string) (*models.User, error) {
	u, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	pass, err := encryption.FromBase64(u.Password)
	if err != nil {
		return nil, err
	}
	decryptedPassword, err := encryption.Decrypt(pass)
	if err != nil {
		return nil, err
	}
	if password != string(decryptedPassword) {
		return nil, ErrInvalidCredentials
	}

	return &models.User{
		ID:    u.ID,
		Email: u.Email,
		Name:  u.Name,
	}, nil
}
