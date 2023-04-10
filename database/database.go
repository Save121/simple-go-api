package database

import (
	"fmt"
	"github.com/Save121/simple-go-api/settings"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func New(s *settings.Settings) (*sqlx.DB, error) {
	connectionString := fmt.Sprintf(
	"user=%s password=%s host=%s port=%d dbname=%s sslmode=disable", 
	s.DB.User, 
	s.DB.Password, 
	s.DB.Host, 
	s.DB.Port, 
	s.DB.Name,
	)
	 return sqlx.Connect("postgres", connectionString) 
}
