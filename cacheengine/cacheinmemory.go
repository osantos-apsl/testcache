package cacheengine

import (
	"sync"
)

type CacheInMemory struct {
	CachedMap *sync.Map
}

func (c *CacheInMemory) Get(key *uint64) *CachedItem {
	var ret *CachedItem = nil
	val, ok := c.CachedMap.Load(*key)
	if ok {
		return val.(*CachedItem)
	}
	return ret
}

func (c *CacheInMemory) Set(key *uint64, pValue *CachedItem) (bool, error) {
	c.CachedMap.Store(*key, pValue)
	return true, nil
}

func (c *CacheInMemory) Dump() []*CachedItem {
	var lAllData []*CachedItem
	c.CachedMap.Range(
		func(key, value any) bool {
			lAllData = append(lAllData, value.(*CachedItem))
			return true
		})
	return lAllData
}

// Secure by cacheEngine once method
func CreateInMemory() *I_Cache {
	var i_c I_Cache
	c := new(CacheInMemory)
	c.CachedMap = new(sync.Map)
	i_c = c
	return &i_c
}
