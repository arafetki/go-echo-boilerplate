package services

import (
	"github.com/arafetki/go-echo-boilerplate/internal/db"
	"github.com/arafetki/go-echo-boilerplate/internal/db/sqlc"
)

type Service struct {
	Users interface {
		Create(params sqlc.InsertUserParams) error
	}
}

func New(db *db.DB) *Service {
	q := sqlc.New(db)
	return &Service{
		Users: &usersService{q},
	}
}
