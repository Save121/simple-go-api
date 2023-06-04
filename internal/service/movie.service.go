package service

import (
	context "context"
	"errors"

	models "github.com/Save121/simple-go-api/internal/models"
)

var (
	validRolesToAddProduct = []string{"d440107f-9d8d-4611-928c-987e9ed05488", "543ff79a-fa2b-4885-ab7c-b7a30e7e60ec"}
	ErrInvalidPermissions  = errors.New("user does not have permission to add movie")
)

func (s *serv) GetMovies(ctx context.Context) ([]models.Movie, error) {
	mm, err := s.repo.GetMovies(ctx)
	if err != nil {
		return nil, err
	}
	movies := []models.Movie{}
	for _, m := range mm {
		movies = append(movies, models.Movie{
			ID:            m.ID,
			Name:          m.Name,
			Description:   m.Description,
			Price:         m.Price,
			Creation_date: m.Creation_date,
		})
	}
	return movies, nil

}
func (s *serv) GetMovieByID(ctx context.Context, id string) (*models.Movie, error) {
	m, err := s.repo.GetMovieByID(ctx, id)
	if err != nil {
		return nil, err
	}
	movie := &models.Movie{
		ID:            m.ID,
		Name:          m.Name,
		Description:   m.Description,
		Price:         m.Price,
		Creation_date: m.Creation_date,
	}
	return movie, nil
}

func (s *serv) AddMovie(ctx context.Context, movie models.Movie, email string) error {
	u, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return err
	}
	roles, err := s.repo.GetUserRoles(ctx, u.ID)
	if err != nil {
		return err
	}
	if len(roles) == 0 {
		return ErrInvalidPermissions
	}
	userCanAdd := false

	for _, r := range roles {
		for _, vr := range validRolesToAddProduct {
			if vr == r.RoleID {
				userCanAdd = true
			}
		}
	}
	if !userCanAdd {
		return ErrInvalidPermissions
	}
	return s.repo.SaveMovie(ctx, movie.Price, movie.Name, movie.Description, u.ID)
}
