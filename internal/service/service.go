package service

import (
	"context"

	"github.com/Save121/simple-go-api/internal/models"
	"github.com/Save121/simple-go-api/internal/repository"
)

// Service is the logic of the application
//
//go:generate mockery --name=Service  --output=service --inpackage
type Service interface {
	RegisterUser(ctx context.Context, email, name, password string) error
	LoginUser(ctx context.Context, email, password string) (*models.User, error)
	AddUserRole(ctx context.Context, userID, roleID string) error
	RemoveUserRole(ctx context.Context, userID, roleID string) error
	GetMovies(ctx context.Context) ([]models.Movie, error)
	GetMovieByID(ctx context.Context, id string) (*models.Movie, error)
	AddMovie(ctx context.Context, movie models.Movie, email string) error
}
type serv struct {
	repo repository.Repository
}

func New(repo repository.Repository) Service {
	return &serv{
		repo: repo,
	}
}
