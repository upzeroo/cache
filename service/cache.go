package service

import (
	"time"

	"github.com/sirupsen/logrus"
	"github.com/upzeroo/cache"
)

type (
	CacheService struct {
		adapter cache.Adapter
		logger  *logrus.Entry
	}
)

func NewCacheService(adapter cache.Adapter, logger *logrus.Entry) *CacheService {
	return &CacheService{
		adapter: adapter,
		logger:  logger,
	}
}

func (serv *CacheService) Get(key string) (string, error) {
	return serv.adapter.Get(key)
}

func (serv *CacheService) Set(key string, data interface{}, exp time.Duration) error {
	return serv.adapter.Set(key, data, exp)
}

func (serv *CacheService) Delete(key string) error {
	return serv.adapter.Delete(key)
}
