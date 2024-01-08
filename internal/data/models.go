package data

import "github.com/jmoiron/sqlx"

type Models struct {
	Users UserModel
}

func NewModels(db *sqlx.DB) Models {
	return Models{
		Users: UserModel{db},
	}
}
