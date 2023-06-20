package adapters

import (
	"fmt"
	"time"
)

type (
	AbstractAdapter struct {
	}
)

var (
	ErrMethodNotImplemented = fmt.Errorf("method not implemented")
)

func NewAbstractAdapter() *AbstractAdapter {
	return &AbstractAdapter{
		// ...
	}
}

func (adapter *AbstractAdapter) Get(key string) (string, error) {
	return "", ErrMethodNotImplemented
}

func (adapter *AbstractAdapter) Set(key string, data interface{}, exp time.Duration) error {
	return ErrMethodNotImplemented
}

func (adapter *AbstractAdapter) Delete(key string) error {
	return ErrMethodNotImplemented
}
