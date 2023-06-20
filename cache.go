package cache

import "time"

type (
	Adapter interface {
		Get(key string) (string, error)
		Set(key string, data interface{}, exp time.Duration) error
		Delete(key string) error
	}
)
