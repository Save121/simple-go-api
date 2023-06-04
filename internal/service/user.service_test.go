package service

import (
	"context"
	"testing"
)

func TestRegisterUser(t *testing.T) {
	testCases := []struct {
		Name          string
		Email         string
		UserName      string
		Password      string
		ExpectedError error
	}{{
		Name:          "RegisterUser_Success",
		Email:         "test@test.com",
		UserName:      "test",
		Password:      "validPassword",
		ExpectedError: nil,
	},
		{
			Name:          "RegisterUser_UserAlreadyExists",
			Email:         "test@exists.com",
			UserName:      "test",
			Password:      "validPassword",
			ExpectedError: ErrUserAlreadyExists,
		}}
	ctx := context.Background()

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			repo.Mock.Test(t)
			err := s.RegisterUser(ctx, tc.Email, tc.UserName, tc.Password)

			if err != tc.ExpectedError {
				t.Errorf("Expected error %v, got %v", tc.ExpectedError, err)
			}
		})

	}
}

func TestLoginUser(t *testing.T) {
	testCases := []struct {
		Name          string
		Email         string
		Password      string
		ExpectedError error
	}{{
		Name:          "LoginUser_Success",
		Email:         "test@exists.com",
		Password:      "validPassword",
		ExpectedError: nil,
	}, {
		Name:          "LoginUser_InvalidPassword",
		Email:         "test@exists.com",
		Password:      "invalidPassword",
		ExpectedError: ErrInvalidCredentials,
	}}

	ctx := context.Background()

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			repo.Mock.Test(t)
			_, err := s.LoginUser(ctx, tc.Email, tc.Password)
			if err != tc.ExpectedError {
				t.Errorf("Expected error %v, got %v", tc.ExpectedError, err)
			}
		})
	}
}
func TestAddUserRole(t *testing.T) {
	testCases := []struct {
		Name          string
		UserID        string
		RoleID        string
		ExpectedError error
	}{{
		Name:          "AddUserRole_Success",
		UserID:        "user-uuid2",
		RoleID:        "role-id2",
		ExpectedError: nil,
	}, {
		Name:          "userAlreadyHasRole",
		UserID:        "user-uuid",
		RoleID:        "role-uuid",
		ExpectedError: ErrUserRoleFound,
	}}

	ctx := context.Background()
	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			repo.Mock.Test(t)
			err := s.AddUserRole(ctx, tc.UserID, tc.RoleID)
			if err != tc.ExpectedError {
				t.Errorf("Expected error %v, got %v", tc.ExpectedError, err)
			}
		})

	}
}

func TestRemoveUserRole(t *testing.T) {
	testCases := []struct {
		Name          string
		UserID        string
		RoleID        string
		ExpectedError error
	}{
		{
			Name:          "RemoveUserRole_Success",
			UserID:        "user-uuid",
			RoleID:        "role-uuid",
			ExpectedError: nil,
		},
	}

	ctx := context.Background()
	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			repo.Mock.Test(t)
			err := s.RemoveUserRole(ctx, tc.UserID, tc.RoleID)
			if err != tc.ExpectedError {
				t.Errorf("Expected error %v, got %v", tc.ExpectedError, err)
			}
		})

	}
}
