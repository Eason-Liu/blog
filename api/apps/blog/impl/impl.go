package impl

import (
	"blog/conf"
	"database/sql"
)

func NewImpl() *Impl {
	return &Impl{}
}

type Impl struct {
	db *sql.DB
}

func (i *Impl) Init() error {
	i.db = conf.Conf().Mysql.GetDB()
	return nil
}
