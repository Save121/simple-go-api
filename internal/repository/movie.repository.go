package repository

import (
	context "context"

	entity "github.com/Save121/simple-go-api/internal/entity"
)

const (
	queryInsertMovie  = "INSERT INTO MOVIES (name, description, price, created_by) values (?, ?, ?, ?)"
	queryGetMovies    = "SELECT id, name, description, price, created_by, creation_date FROM MOVIES"
	queryGetMovieByID = "SELECT id, name, description, price, created_by, creation_date FROM MOVIES WHERE id = ?"
)

func (r *repo) SaveMovie(ctx context.Context, price float32, name, description, createdBy string) error {
	_, err := r.db.ExecContext(ctx, queryInsertMovie, name, description, price, createdBy)
	return err
}

func (r *repo) GetMovies(ctx context.Context) ([]entity.Movie, error) {
	movies := []entity.Movie{}
	err := r.db.SelectContext(ctx, &movies, queryGetMovies)
	if err != nil {
		return nil, err
	}
	return movies, nil
}

func (r *repo) GetMovieByID(ctx context.Context, id string) (*entity.Movie, error) {
	movie := &entity.Movie{}
	err := r.db.GetContext(ctx, movie, queryGetMovies, id)
	if err != nil {
		return nil, err
	}
	return movie, nil
}
