package adapters

import (
	"errors"

	"github.com/upzeroo/cache"
)

type (
	DepContainer struct {
		RedisURL string
	}
)

var (
	ErrAdapterNotSupported = errors.New("adapter not supported")
)

const (
	Redis = "redis"
)

func Factory(providerName string, dep *DepContainer) (cache.Adapter, error) {
	switch providerName {
	case Redis:
		return NewRedisAdapter(dep)
	default:
		return nil, ErrAdapterNotSupported
	}
}
