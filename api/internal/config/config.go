package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	DataBase DataBase
	Auth     Auth
}

type Auth struct {
	AccessSecret string
	AccessExpire int64
}
type DataBase struct {
	DataSource string
}
