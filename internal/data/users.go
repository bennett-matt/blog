package data

import "github.com/jmoiron/sqlx"

type UserModel struct {
	db *sqlx.DB
}

func (um UserModel) Insert() {

}
