package repository

import (
	"context"

	"github.com/Save121/simple-go-api/internal/entity"
)

const (
	queryInsertUser = `
	INSERT INTO USERS (email, name, password) 
	VALUES($1, $2, $3)`

	queryGetUserByEmail = `
	SELECT 
	id, 
	email, 
	name, 
	password 
	FROM USERS 
	WHERE email = $1`
)

func (r *repo) SaveUser(ctx context.Context, email, name, password string) error {
	_, err := r.db.ExecContext(ctx, queryInsertUser, email, name, password)
	return err
}
func (r *repo) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	u := &entity.User{}
	err := r.db.GetContext(ctx, u, queryGetUserByEmail, email)
	if err != nil {
		return nil, err
	}

	return u, nil
}
