package dao

import (
	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/jinzhu/gorm"
	"go-common/library/database/orm"
)

func NewDB() (db *gorm.DB, cf func(), err error) {
	var (
		cfg orm.Config
		ct paladin.TOML
	)
	if err = paladin.Get("db.toml").Unmarshal(&ct); err != nil {
		return
	}
	if err = ct.Get("Client").UnmarshalTOML(&cfg); err != nil {
		return
	}
	db = orm.NewMysql(&cfg)
	cf = func() {db.Close()}
	return
}
