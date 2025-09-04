package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	DataSource string
	Auth    JwtAuth
}

type JwtAuth struct {
	AccessSecret string
	AccessExpire int64
}
