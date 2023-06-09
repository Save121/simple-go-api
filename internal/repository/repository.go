package repository

import (
	"context"

	"github.com/Save121/simple-go-api/internal/entity"
	"github.com/jmoiron/sqlx"
)

// Repository is the interface that wraps tha basic CRUD operations
//
//go:generate mockery --name=Repository  --output=repository --inpackage
type Repository interface {
	SaveUser(ctx context.Context, email, name, password string) error
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
	SaveUserRole(ctx context.Context, userID, roleID string) error
	RemoveUserRole(ctx context.Context, userID, roleID string) error
	GetUserRoles(ctx context.Context, userID string) ([]entity.UserRole, error)
	SaveMovie(ctx context.Context, price float32, name, description, createdBy string) error
	GetMovies(ctx context.Context) ([]entity.Movie, error)
	GetMovieByID(ctx context.Context, id string) (*entity.Movie, error)
}

type repo struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) Repository {
	return &repo{
		db: db,
	}
}
