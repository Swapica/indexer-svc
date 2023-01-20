package config

import (
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/kit/pgdb"
)

type Config interface {
	comfig.Logger
	pgdb.Databaser

	Network() Network
}

type config struct {
	comfig.Logger
	pgdb.Databaser
	getter kv.Getter
	once   comfig.Once
}

func New(getter kv.Getter) Config {
	return &config{
		getter:    getter,
		Databaser: pgdb.NewDatabaser(getter),
		Logger:    comfig.NewLogger(getter, comfig.LoggerOpts{}),
	}
}
