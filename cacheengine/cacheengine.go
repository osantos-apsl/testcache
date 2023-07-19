package cacheengine

import (
	"errors"
	"sync"
	"time"
)

type CacheManager struct {
	layers []*I_Cache
}

type CachedItem struct {
	HashID   *uint64
	Data     []byte
	LastDate time.Time
	Error    error
}

type I_Cache interface {
	Get(key *uint64) *CachedItem
	Set(key *uint64, value *CachedItem) (bool, error)
	Dump() []*CachedItem
}

var once sync.Once
var layers sync.Once

func (c *CacheManager) RegisterLayer(createManager func() *I_Cache) {
	layers.Do(func() {
		c.layers = append(c.layers, createManager())
	})
}

func (c *CacheManager) Get(key *uint64, pLvl ...int) *CachedItem {
	if len(c.layers) == 0 {
		return nil
	}
	if len(pLvl) == 0 {
		return (*c.layers[0]).Get(key)

	}
	// for leveling cached
	// for i := range pLvl {
	// 	ret := (*c.layers[i]).Get(key)
	// 	if ret != nil {
	// 		return ret
	// 	}

	// }
	return nil
}

func (c *CacheManager) Set(key *uint64, pValue *CachedItem, pLvl ...int) (bool, error) {
	if len(c.layers) == 0 {
		return false, errors.New("No cache layer found!!")
	}
	if len(pLvl) > 1 {
		return false, errors.New("Lvl expected!!")
	}
	(*pValue).HashID = key
	if len(pLvl) == 0 {
		return (*c.layers[0]).Set(key, pValue)
	} else {
		return (*c.layers[pLvl[0]]).Set(key, pValue)
	}
}

func (c *CacheManager) Dump() []*CachedItem {
	if len(c.layers) == 0 {
		return nil
	}
	return (*c.layers[0]).Dump()
}

var cache *CacheManager = nil

func GetCacheManager() *CacheManager {
	once.Do(func() {
		cache = new(CacheManager)
		cache.layers = make([]*I_Cache, 0)
	})
	//log.Println("Manager Mem Addr: ", &cache)
	return cache
}
