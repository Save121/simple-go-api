package service

import (
	"os"
	"testing"

	"github.com/Save121/simple-go-api/encryption"
	"github.com/Save121/simple-go-api/internal/entity"
	"github.com/Save121/simple-go-api/internal/repository"
	mock "github.com/stretchr/testify/mock"
)

var repo *repository.MockRepository
var s Service

func TestMain(m *testing.M) {
	validPassword, _ := encryption.Encrypt([]byte("validPassword"))
	encryptedPassword := encryption.ToBase64(validPassword)
	u := &entity.User{Email: "test@exists.com", Password: encryptedPassword}
	adminUser := &entity.User{ID: "uuid-1", Email: "admin@email.com", Password: encryptedPassword}
	customerUser := &entity.User{ID: "uuid-2", Email: "customer@email.com", Password: encryptedPassword}
	repo = &repository.MockRepository{}
	repo.On("GetUserByEmail", mock.Anything, "test@test.com").Return(nil, nil)
	repo.On("GetUserByEmail", mock.Anything, "test@exists.com").Return(u, nil)
	repo.On("GetUserByEmail", mock.Anything, "admin@email.com").Return(adminUser, nil)
	repo.On("GetUserByEmail", mock.Anything, "customer@email.com").Return(customerUser, nil)

	repo.On("SaveUser", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)
	repo.On("SaveUserRole", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	repo.On("GetUserRoles", mock.Anything, string("user-uuid")).Return([]entity.UserRole{{
		UserID: "user-uuid",
		RoleID: "role-uuid",
	}}, nil)
	repo.On("GetUserRoles", mock.Anything, string("uuid-1")).Return([]entity.UserRole{{UserID: "uuid-1", RoleID: "d440107f-9d8d-4611-928c-987e9ed05488"}}, nil)
	repo.On("GetUserRoles", mock.Anything, string("uuid-2")).Return([]entity.UserRole{{UserID: "uuid-2", RoleID: "invalid-uuid-to-add-movie"}}, nil)
	repo.On("RemoveUserRole", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	repo.On("GetUserRoles", mock.Anything, mock.Anything).Return(nil, nil)

	repo.On("SaveMovie", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)

	s = New(repo)
	code := m.Run()
	os.Exit(code)
}
