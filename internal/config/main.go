package config

import (
	jsonapi "gitlab.com/distributed_lab/json-api-connector"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
)

type Config interface {
	comfig.Logger

	Network() Network
	Collector() *jsonapi.Connector
}

type config struct {
	comfig.Logger
	getter kv.Getter

	networkOnce   comfig.Once
	collectorOnce comfig.Once
}

func New(getter kv.Getter) Config {
	return &config{
		getter: getter,
		Logger: comfig.NewLogger(getter, comfig.LoggerOpts{}),
	}
}
