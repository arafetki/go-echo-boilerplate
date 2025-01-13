package services

import "github.com/arafetki/go-echo-boilerplate/internal/db/sqlc"

type usersService struct {
	q *sqlc.Queries
}

func (us *usersService) Create(params sqlc.InsertUserParams) error {
	return nil
}
