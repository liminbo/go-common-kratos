package dao

import (
	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"go-common/library/database/es"
)

func NewES() (elasticSearch *es.ElasticSearch, cf func(), err error) {
	var (
		cfg es.Config
		ct paladin.Map
	)
	if err = paladin.Get("es.toml").Unmarshal(&ct); err != nil {
		return
	}
	if err = ct.Get("Client").UnmarshalTOML(&cfg); err != nil {
		return
	}
	elasticSearch = es.New(&cfg)
	cf = func() {}
	return
}
