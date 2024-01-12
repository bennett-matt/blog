package data

import (
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
)

type Models struct {
	Users  UserModel
	Tokens TokenModel
}

func NewModels(db *pgxpool.Pool) Models {
	return Models{
		Users:  UserModel{db},
		Tokens: TokenModel{db},
	}
}
