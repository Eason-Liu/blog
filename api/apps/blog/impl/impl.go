package impl

import (
	"blog/conf"
	"gorm.io/gorm"
)

func NewImpl() *Impl {
	return &Impl{}
}

type Impl struct {
	db *gorm.DB
}

func (i *Impl) Name() string {
	return "blog"
}

func (i *Impl) DB() *gorm.DB {
	return i.db.Table(i.Name())
}

func (i *Impl) Init() error {
	i.db = conf.Conf().Mysql.GetORMDB().Debug()
	return nil
}
