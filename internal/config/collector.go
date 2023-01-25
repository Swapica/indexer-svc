package config

import (
	"net/http"
	"net/url"
	"time"

	"gitlab.com/distributed_lab/figure/v3"
	jsonapi "gitlab.com/distributed_lab/json-api-connector"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/connectors/signed"
)

func (c *config) Collector() *jsonapi.Connector {
	return c.collectorOnce.Do(func() interface{} {
		var cfg struct {
			Endpoint       *url.URL      `fig:"endpoint,required"`
			RequestTimeout time.Duration `fig:"request_timeout"`
		}
		err := figure.Out(&cfg).
			From(kv.MustGetStringMap(c.getter, "collector")).
			Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out collector"))
		}

		if cfg.RequestTimeout == 0 {
			cfg.RequestTimeout = defaultRequestTimeout
		}

		return jsonapi.NewConnector(signed.NewClient(&http.Client{Timeout: cfg.RequestTimeout}, cfg.Endpoint))
	}).(*jsonapi.Connector)
}
