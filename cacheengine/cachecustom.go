package cacheengine

import "sync"

type CacheCustomInMemory struct {
	CachedMap map[uint64]*CachedItem
	lock      sync.RWMutex
}

func (c *CacheCustomInMemory) Get(key *uint64) *CachedItem {
	c.lock.RLock()
	val, isPresent := c.CachedMap[*key]
	c.lock.RUnlock()
	if isPresent {
		return val
	}
	return nil
}

func (c *CacheCustomInMemory) Set(key *uint64, pValue *CachedItem) (bool, error) {
	c.lock.Lock()
	c.CachedMap[*key] = pValue
	//[*key] = pValue
	c.lock.Unlock()
	return true, nil
}

func (c *CacheCustomInMemory) Dump() []*CachedItem {
	var lAllData []*CachedItem
	for _, value := range c.CachedMap {
		lAllData = append(lAllData, value)
	}
	return lAllData
}

// Secure by cacheEngine once method
func CreateCustom() *I_Cache {
	var i_c I_Cache
	c := new(CacheCustomInMemory)
	c.CachedMap = make(map[uint64]*CachedItem)
	i_c = c
	return &i_c
}
