package dao

import (
	"context"

	"kvm-agent/internal/dal/cache"
	"kvm-agent/internal/dal/gen"
)

// Dao dao top level.
type Dao struct {
	ctx   context.Context
	query *gen.Query
	cache *cache.Cache
}

// GetQuery get query.
func (d *Dao) GetQuery() *gen.Query {
	return d.query
}

// GetCtx get context.
func (d *Dao) GetCtx() context.Context {
	return d.ctx
}

// GetCache get cache.
func (d *Dao) GetCache() *cache.Cache {
	return d.cache
}
