package impl

import (
	"database/sql"
)

func NewImpl() *Impl {
	return &Impl{}
}

type Impl struct {
	db *sql.DB
}

func (i *Impl) init() error {
	i.db =
	return nil
}
