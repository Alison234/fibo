package cache

import (
	"bytes"
	"encoding/gob"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/go/http-rest-api/fibonacci"
)

type Cache interface {
	GetValue(key string) ([]fibonacci.Fibonacci, error)
	SetValue(key string, value []fibonacci.Fibonacci) error
}

type MemCacher struct {
	addr       string
	cacheStore *memcache.Client
}

func NewMemCacher(addr string) *MemCacher {
	var client = memcache.New(addr)
	client.MaxIdleConns = 5
	client.Timeout = 1 * time.Minute
	return &MemCacher{addr: addr, cacheStore: client}
}

func (m *MemCacher) GetValue(key string) ([]fibonacci.Fibonacci, error) {
	item, err := m.cacheStore.Get(key)
	if err != nil {
		return nil, err
	}
	b := bytes.NewReader(item.Value)

	var res []fibonacci.Fibonacci

	if err := gob.NewDecoder(b).Decode(&res); err != nil {
		return nil, err
	}

	return res, nil
}

func (m *MemCacher) SetValue(key string, value []fibonacci.Fibonacci) error {
	var b bytes.Buffer

	if err := gob.NewEncoder(&b).Encode(value); err != nil {
		return err
	}

	return m.cacheStore.Set(&memcache.Item{
		Key:   key,
		Value: b.Bytes(),
	})
}
