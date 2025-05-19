package sqlite

import (
	"database/sql"
	_ "embed"
	"fmt"

	"git.fruzit.pp.ua/weather/api/pkg/user/domain/entity"
	"git.fruzit.pp.ua/weather/api/pkg/user/port"
)

//go:embed bob/schema.sql
var Schema []byte

type Sqlite struct {
	db *sql.DB
}

var _ port.Storage = (*Sqlite)(nil)

func NewSqlite(db *sql.DB) *Sqlite {
	return &Sqlite{db}
}

func (s *Sqlite) SaveUser(user entity.User) error {
	// TODO: SaveUser
	return fmt.Errorf("not implemented")
}
