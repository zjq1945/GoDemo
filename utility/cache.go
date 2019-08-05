package utility

import (
	//"fmt"
	"github.com/patrickmn/go-cache"
	"time"
)

var myCache *cache.Cache

func InitCache() {
	myCache = cache.New(999*time.Hour, 999*time.Hour)
}

func AddCache(key string, value interface{}, cacheDuration time.Duration) {

	myCache.Set(key, value, cacheDuration)
}

func GetCache(key string) (interface{}, bool) {
	rtn, result := myCache.Get(key)
	return rtn, result
}
