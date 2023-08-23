package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
)

var _ Cache = new(cache)

// Cache cache interface.
type Cache interface {
	// default
	Set(ctx context.Context, key, value string, ttl time.Duration) error
	Get(ctx context.Context, key string) (string, error)
	TTL(ctx context.Context, key string) (time.Duration, error)
	Expire(ctx context.Context, key string, ttl time.Duration) bool
	ExpireAt(ctx context.Context, key string, ttl time.Time) bool
	Del(ctx context.Context, key string) bool
	Exists(ctx context.Context, keys ...string) bool
	Incr(ctx context.Context, key string) int64
	Close() error

	// List
	LPush(ctx context.Context, key string, values ...interface{}) error
	RPush(ctx context.Context, key string, values ...interface{}) error
	LPop(ctx context.Context, key string) (string, error)
	RPop(ctx context.Context, key string) (string, error)
	LRange(ctx context.Context, key string, start, stop int64) ([]string, error)
	LRem(ctx context.Context, key string, count int64, value interface{}) (int64, error)
	LTrim(ctx context.Context, key string, start, stop int64) error
	LLen(ctx context.Context, key string) (int64, error)

	// ZSet
	ZAdd(ctx context.Context, key string, members ...*redis.Z) (int64, error)
	ZRem(ctx context.Context, key string, members ...interface{}) (int64, error)
	ZRange(ctx context.Context, key string, start, stop int64) ([]string, error)

	// Map
	HSet(ctx context.Context, key, field string, value string) error
	HGet(ctx context.Context, key, field string) (string, error)
	HGetAll(ctx context.Context, key string) (map[string]string, error)
	HDel(ctx context.Context, key string, fields ...string) (int64, error)
	HExists(ctx context.Context, key, field string) (bool, error)
	HIncrBy(ctx context.Context, key, field string, incr int64) (int64, error)
	HKeys(ctx context.Context, key string) ([]string, error)
	HLen(ctx context.Context, key string) (int64, error)
	HMGet(ctx context.Context, key string, fields ...string) ([]interface{}, error)
	HMSet(ctx context.Context, key string, values ...interface{}) error

	// tx
	Pipelined(ctx context.Context, fn func(redis.Pipeliner) error) ([]redis.Cmder, error)
	Pipline() redis.Pipeliner
	TxPipelined(ctx context.Context, fn func(redis.Pipeliner) error) ([]redis.Cmder, error)
	TxPipeline() redis.Pipeliner

	Watch(ctx context.Context, fn func(*redis.Tx) error, keys ...string) error
}

// cache cache dao representative.
type cache struct {
	client *redis.Client
}

// Use return cache db.
func Use(client *redis.Client) Cache {
	return &cache{
		client: client,
	}
}

// Get get value.
func (c *cache) Get(ctx context.Context, key string) (string, error) {
	value, err := c.client.Get(ctx, key).Result()
	if err != nil {
		return "", errors.Wrapf(err, "redis get key: %s err", key)
	}

	return value, nil
}

// Set set key value pair to redis.
func (c *cache) Set(ctx context.Context, key string, value string, ttl time.Duration) error {
	if err := c.client.Set(ctx, key, value, ttl).Err(); err != nil {
		return errors.Wrapf(err, "redis set key: %s err", key)
	}

	return nil
}

// TTL get ttl of key.
func (c *cache) TTL(ctx context.Context, key string) (time.Duration, error) {
	ttl, err := c.client.TTL(ctx, key).Result()
	if err != nil {
		return -1, errors.Wrapf(err, "redis get key: %s err", key)
	}

	return ttl, nil
}

// Expire expire key.
func (c *cache) Expire(ctx context.Context, key string, ttl time.Duration) bool {
	ok, _ := c.client.Expire(ctx, key, ttl).Result()

	return ok
}

// ExpireAt expire key at time.
func (c *cache) ExpireAt(ctx context.Context, key string, ttl time.Time) bool {
	ok, _ := c.client.ExpireAt(ctx, key, ttl).Result()

	return ok
}

// Exists check some keys are exist.
func (c *cache) Exists(ctx context.Context, keys ...string) bool {
	if len(keys) == 0 {
		return true
	}
	value, _ := c.client.Exists(ctx, keys...).Result()

	return value > 0
}

// Del delete key.
func (c *cache) Del(ctx context.Context, key string) bool {
	if key == "" {
		return true
	}

	value, _ := c.client.Del(ctx, key).Result()

	return value > 0
}

// Incr incr key.
func (c *cache) Incr(ctx context.Context, key string) int64 {
	value, _ := c.client.Incr(ctx, key).Result()

	return value
}

// Close close redis client.
func (c *cache) Close() error {
	return c.client.Close()
}

// LPush left push value to list.
func (c *cache) LPush(ctx context.Context, key string, values ...interface{}) error {
	return c.client.LPush(ctx, key, values...).Err()
}

// RPush right push value to list.
func (c *cache) RPush(ctx context.Context, key string, values ...interface{}) error {
	return c.client.RPush(ctx, key, values...).Err()
}

// LPop left pop value from list.
func (c *cache) LPop(ctx context.Context, key string) (string, error) {
	return c.client.LPop(ctx, key).Result()
}

// RPop right pop value from list.
func (c *cache) RPop(ctx context.Context, key string) (string, error) {
	return c.client.RPop(ctx, key).Result()
}

// LRange get list value from start to stop.
func (c *cache) LRange(ctx context.Context, key string, start, stop int64) ([]string, error) {
	return c.client.LRange(ctx, key, start, stop).Result()
}

// LRem remove count value from list.
func (c *cache) LRem(ctx context.Context, key string, count int64, value interface{}) (int64, error) {
	return c.client.LRem(ctx, key, count, value).Result()
}

// LTrim trim list from start to stop.
func (c *cache) LTrim(ctx context.Context, key string, start, stop int64) error {
	return c.client.LTrim(ctx, key, start, stop).Err()
}

// LLen get list length.
func (c *cache) LLen(ctx context.Context, key string) (int64, error) {
	return c.client.LLen(ctx, key).Result()
}

// Pipelined pipelined.
func (c *cache) Pipelined(ctx context.Context, fn func(redis.Pipeliner) error) ([]redis.Cmder, error) {
	return c.client.Pipelined(ctx, fn)
}

// Pipline pipline.
func (c *cache) Pipline() redis.Pipeliner {
	return c.client.Pipeline()
}

// TxPipelined tx pipelined.
func (c *cache) TxPipelined(ctx context.Context, fn func(redis.Pipeliner) error) ([]redis.Cmder, error) {
	return c.client.TxPipelined(ctx, fn)
}

// TxPipeline tx pipeline.
func (c *cache) TxPipeline() redis.Pipeliner {
	return c.client.TxPipeline()
}

// Watch watch.
func (c *cache) Watch(ctx context.Context, fn func(*redis.Tx) error, keys ...string) error {
	return c.client.Watch(ctx, fn, keys...)
}

// LPopAndRPush lpop and rpush.
func (c *cache) LPopAndRPush(ctx context.Context, key string) (string, error) {
	tx := c.client.TxPipeline()
	defer tx.Close()

	// check length
	length, err := c.client.LLen(ctx, key).Result()
	if err != nil || length <= 1 {
		return "", errors.Wrapf(err, "list is not exists or length less than 1: %s err", key)
	}

	// get head value
	headVal, _ := c.client.LIndex(ctx, key, 0).Result()

	_ = tx.Process(ctx, c.client.LPop(ctx, key))
	_ = tx.Process(ctx, c.client.RPush(ctx, key, headVal))

	_, err = tx.Exec(ctx)
	if err != nil {
		return "", err
	}

	return headVal, nil
}

// ZAdd zset add.
func (c *cache) ZAdd(ctx context.Context, key string, members ...*redis.Z) (int64, error) {
	return c.client.ZAdd(ctx, key, members...).Result()
}

// ZRem zset remove.
func (c *cache) ZRem(ctx context.Context, key string, members ...interface{}) (int64, error) {
	return c.client.ZRem(ctx, key, members...).Result()
}

// ZRange zset range.
func (c *cache) ZRange(ctx context.Context, key string, start, stop int64) ([]string, error) {
	return c.client.ZRange(ctx, key, start, stop).Result()
}

// // EnqueueStringLPush enqueue key and push left.
// func (c *cache) EnqueueStringLPush(ctx context.Context, key, value string) error {
// 	c.mu.Lock()
// 	defer c.mu.Unlock()

// 	err := c.client.LPush(ctx, key, value).Err()
// 	if err != nil {
// 		return errors.Wrapf(err, "redis enqueue key: %s , value: %s, err: %s", key, value, err.Error())
// 	}

// 	return nil
// }

// // DequeueStringRPop dequeue key and pop right.
// func (c *cache) DequeueStringRPop(ctx context.Context, key string) (string, error) {
// 	c.mu.Lock()
// 	defer c.mu.Unlock()

// 	value, err := c.client.RPop(ctx, key).Result()
// 	if err != nil {
// 		return "", errors.Wrapf(err, "redis dequeue key: %s , err: %s", key, err.Error())
// 	}

// 	return value, nil
// }

// func (c *cache) TxPipeline(ctx context.Context) redis.Pipeliner {
// 	return c.client.TxPipeline()
// }

// HSet hset.
func (c *cache) HSet(ctx context.Context, key, field string, value string) error {
	return c.client.HSet(ctx, key, field, value).Err()
}

// HGet hget.
func (c *cache) HGet(ctx context.Context, key, field string) (string, error) {
	return c.client.HGet(ctx, key, field).Result()
}

// HGetAll hgetall.
func (c *cache) HGetAll(ctx context.Context, key string) (map[string]string, error) {
	return c.client.HGetAll(ctx, key).Result()
}

// HDel hdel.
func (c *cache) HDel(ctx context.Context, key string, fields ...string) (int64, error) {
	return c.client.HDel(ctx, key, fields...).Result()
}

// HExists hexists.
func (c *cache) HExists(ctx context.Context, key, field string) (bool, error) {
	return c.client.HExists(ctx, key, field).Result()
}

// HIncrBy hincrby.
func (c *cache) HIncrBy(ctx context.Context, key, field string, incr int64) (int64, error) {
	return c.client.HIncrBy(ctx, key, field, incr).Result()
}

// HKeys hkeys.
func (c *cache) HKeys(ctx context.Context, key string) ([]string, error) {
	return c.client.HKeys(ctx, key).Result()
}

// HLen hlen.
func (c *cache) HLen(ctx context.Context, key string) (int64, error) {
	return c.client.HLen(ctx, key).Result()
}

// HMGet hmget.
func (c *cache) HMGet(ctx context.Context, key string, fields ...string) ([]interface{}, error) {
	return c.client.HMGet(ctx, key, fields...).Result()
}

// HMSet hmset.
func (c *cache) HMSet(ctx context.Context, key string, values ...interface{}) error {
	return c.client.HMSet(ctx, key, values).Err()
}
