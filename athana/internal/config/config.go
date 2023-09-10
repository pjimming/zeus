package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	MySQL MySQL
}

type MySQL struct {
	DSN string
}
