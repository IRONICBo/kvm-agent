package dao

import (
	"context"
	"errors"
	"kvm-agent/internal/conn"
	"kvm-agent/internal/dal/cache"
	"kvm-agent/internal/dal/gen"
	"sync"
	"time"
)

// TriggerDao trigger dao.
type TriggerDao struct {
	Dao

	mu sync.Mutex
}

// NewTriggerDao return a trigger dao.
func NewTriggerDao() *TriggerDao {
	query := gen.Use(conn.GetDMDB())
	cache := cache.Use(conn.GetRedis())

	return &TriggerDao{
		Dao: Dao{
			ctx:   context.Background(),
			query: query,
			cache: &cache,
		},

		mu: sync.Mutex{},
	}
}

// PushListWithRetry push list with retry.
func (d *TriggerDao) PushListWithRetry(key, data string, retry, interval int) error {
	// Push to redis.
	// check list length, if length > 20, wait forever.
	// d.mu.Lock()
	// defer d.mu.Unlock()

	// TODO: Set max retry times.
	if retry > 20 {
		retry = 20
	}
	if retry < 0 {
		retry = 0
		return errors.New("retry times is 0")
	}

	if interval == 0 {
		return errors.New("interval times is 0")
	}

	// check list length, if length > 20, wait forever.
	if length, _ := (*d.cache).LLen(d.ctx, key); length > 20 {
		// d.mu.Unlock()
		time.Sleep(time.Duration(interval) * time.Second)

		d.PushListWithRetry(key, data, retry-1, interval)
		return errors.New("list length is more than 20")
	}

	err := (*d.cache).LPush(d.ctx, key, data)
	if err != nil {
		// d.mu.Unlock()
		d.PushListWithRetry(key, data, retry-1, interval)
		return errors.New("push list error")
	}

	return nil
}
