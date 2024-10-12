package db

import (
	"github.com/goburrow/cache"
	"math"
	"time"
)

// cache key
const (
	uidKey = "uid.%d.%s" //pid,openId
)

var (
	// uid缓存 key:uidKey, value:uid
	uidCache = cache.New(
		cache.WithMaximumSize(math.MaxUint16),
		cache.WithExpireAfterAccess(2*time.Hour),
	)
)
