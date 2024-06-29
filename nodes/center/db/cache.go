package db

import (
	"github.com/goburrow/cache"
	"math"
	"time"
)

var (

	// uid缓存 key:uidKey, value:uid
	uidCache = cache.New(
		cache.WithMaximumSize(math.MaxUint16),
		cache.WithExpireAfterAccess(2*time.Hour),
	)

	// 开发帐号缓存 key:accountName, value:DevAccountTable
	devAccountCache = cache.New(
		cache.WithMaximumSize(math.MaxUint16),
		cache.WithExpireAfterAccess(time.Hour),
	)
)

// cache key
const (
	uidKey = "uid.%d.%s" //pid,openId
)
