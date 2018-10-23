/**
缓存最新数据
*/
package memcache

import (
	"github.com/coocood/freecache"
	"runtime/debug"
)

var cache *freecache.Cache

func init() {
	//10M
	cacheSize := 10 * 1024 * 1024
	cache = freecache.NewCache(cacheSize)
	debug.SetGCPercent(10)
}


