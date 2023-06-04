package service

import (
	context "context"
	"testing"

	"github.com/Save121/simple-go-api/internal/models"
)

func TestAddMovie(t *testing.T) {
	testCases := []struct {
		Name          string
		Movie         models.Movie
		Email         string
		ExpectedError error
	}{
		{
			Name: "AddMovie_Success",
			Movie: models.Movie{
				Name:        "Test Movie",
				Description: "Test Description",
				Price:       10.00,
			},
			Email:         "admin@email.com",
			ExpectedError: nil,
		},
		{
			Name: "AddMovie_InvalidPermissions",
			Movie: models.Movie{
				Name:        "Test Movie",
				Description: "Test Description",
				Price:       10.00,
			},
			Email:         "customer@email.com",
			ExpectedError: ErrInvalidPermissions,
		},
	}

	ctx := context.Background()

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			repo.Mock.Test(t)

			err := s.AddMovie(ctx, tc.Movie, tc.Email)

			if err != tc.ExpectedError {
				t.Errorf("Expected error %v, got %v", tc.ExpectedError, err)
			}
		})
	}
}
