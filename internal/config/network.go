package config

import (
	"math"
	"strconv"
	"time"

	"github.com/Swapica/indexer-svc/internal/gobind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"gitlab.com/distributed_lab/figure/v3"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type Network struct {
	*gobind.Swapica
	ChainID        string
	IndexPeriod    time.Duration
	RequestTimeout time.Duration
}

const defaultRequestTimeout = 10 * time.Second
const maxChainID int64 = math.MaxUint64/2 - 36

func (c *config) Network() Network {
	return c.networkOnce.Do(func() interface{} {
		var cfg struct {
			RPC            string         `fig:"rpc,required"`
			Contract       common.Address `fig:"contract,required"`
			ChainID        int64          `fig:"chain_id,required"`
			IndexPeriod    time.Duration  `fig:"index_period,required"`
			RequestTimeout time.Duration  `fig:"request_timeout"`
		}

		err := figure.Out(&cfg).
			With(figure.EthereumHooks).
			From(kv.MustGetStringMap(c.getter, "network")).
			Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out network"))
		}

		if cfg.ChainID > maxChainID || cfg.ChainID <= 0 {
			panic("chain_id value out of range due to EIP 2294")
		}
		cli, err := ethclient.Dial(cfg.RPC)
		if err != nil {
			panic(errors.Wrap(err, "failed to connect to RPC provider"))
		}
		s, err := gobind.NewSwapica(cfg.Contract, cli)
		if err != nil {
			panic(errors.Wrap(err, "failed to create contract caller"))
		}

		if cfg.RequestTimeout == 0 {
			cfg.RequestTimeout = defaultRequestTimeout
		}

		return Network{
			Swapica:        s,
			ChainID:        strconv.FormatInt(cfg.ChainID, 10),
			IndexPeriod:    cfg.IndexPeriod,
			RequestTimeout: cfg.RequestTimeout,
		}
	}).(Network)
}
