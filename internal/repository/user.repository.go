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

	queryInsertUserRole = `
	INSERT INTO ROLES_USERS 
	(use_id, role_id) 
	VALUES (:user_id, :role_id)`

	queryRemoveUserRole = `
	DELETE FROM ROLES_USERS
	WHERE user_id = :user_id
	AND role_id = :role_id`
	querySelectUserRoles = `
	SELECT user_id, role_id
	FROM ROLES_USERS 
	WHERE user_id = :user_id`
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

func (r *repo) SaveUserRole(ctx context.Context, userID, roleID string) error {
	data := entity.UserRole{
		UserID: userID,
		RoleID: roleID,
	}
	_, err := r.db.NamedExecContext(ctx, queryInsertUserRole, data)
	return err
}
func (r *repo) RemoveUserRole(ctx context.Context, userID, roleID string) error {
	data := entity.UserRole{
		UserID: userID,
		RoleID: roleID,
	}
	_, err := r.db.NamedExecContext(ctx, queryRemoveUserRole, data)
	return err
}

func (r *repo) GetUserRoles(ctx context.Context, userID string) ([]entity.UserRole, error) {
	roles := []entity.UserRole{}
	if err := r.db.SelectContext(ctx, &roles, querySelectUserRoles, userID); err != nil {
		return nil, err
	}

	return roles, nil
}
