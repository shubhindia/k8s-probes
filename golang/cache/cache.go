package cache

import (
	"time"

	"github.com/jellydator/ttlcache/v3"
)

func GetCache() *ttlcache.Cache[string, string] {
	cache := ttlcache.New[string, string](
		ttlcache.WithTTL[string, string](30 * time.Hour),
	)

	return cache
}
