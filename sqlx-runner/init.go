package runner

import (
	"github.com/mgutz/logxi/v1"
	"gopkg.in/mgutz/dat.v1"
	"gopkg.in/mgutz/dat.v1/kvs"
	"gopkg.in/mgutz/dat.v1/postgres"
)

var logger log.Logger

func init() {
	dat.Dialect = postgres.New()
	logger = log.New("dat:sqlx")
}

// Cache caches query results.
var Cache kvs.KeyValueStore

// SetCache sets this runner's cache. The default cache is in-memory
// based. See cache.MemoryKeyValueStore.
func SetCache(store kvs.KeyValueStore) {
	Cache = store
}
