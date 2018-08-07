package memcache

import (
	"encoding/json"
	"log"
	"time"

	ttlCache "github.com/koding/cache"
)

type IMemCache interface {
	SetCacheTTL(memCacheName string, key string, value interface{}) error
	GetCacheTTL(memCacheName string, key string) (interface{}, error)
	SetCacheTTLJSON(memCacheName string, key string, value interface{}) error
	GetCacheTTLJSON(memCacheName string, key string, value interface{}) error
	DeleteCache(memCacheName string, key string) error
}

const (
	Key15Min = "15m"
	Key30Min = "30m"
)

type KodingCache struct {
	MapMemoryTTL map[string]*ttlCache.MemoryTTL
}

var kodingCacheObj = KodingCache{
	MapMemoryTTL: make(map[string]*ttlCache.MemoryTTL),
}

func InitKodingCache() {
	mapDuration := map[string]time.Duration{
		"15m": 15 * time.Minute,
		"30m": 15 * time.Minute,
	}

	for key, value := range mapDuration {
		kodingCacheObj.MapMemoryTTL[key] = ttlCache.NewMemoryWithTTL(value)
		kodingCacheObj.MapMemoryTTL[key].StartGC(time.Second)
	}
}

func GetKodingCache() *KodingCache {
	return &kodingCacheObj
}

func (c *KodingCache) SetCacheTTL(memCacheName string, key string, value interface{}) error {
	if err := c.MapMemoryTTL[memCacheName].Set(key, value); err != nil {
		return err
	}

	return nil
}

func (c *KodingCache) GetCacheTTL(memCacheName string, key string) (interface{}, error) {
	value, err := c.MapMemoryTTL[memCacheName].Get(key)
	if err != nil {
		return value, err
	}

	return value, nil
}

func (c *KodingCache) SetCacheTTLJSON(memCacheName string, key string, value interface{}) error {
	encoded, err := json.Marshal(value)
	if err != nil {
		return err
	}

	if err := c.SetCacheTTL(memCacheName, key, encoded); err != nil {
		log.Printf("Error SetCacheTTLJSON for MemCache %s, Key %s, Error : {%+v}\n", memCacheName, key, err)
		return err
	}

	log.Printf("Success SetCacheTTLJSON for MemCache %s, Key %s\n", memCacheName, key)
	return nil
}

func (c *KodingCache) GetCacheTTLJSON(memCacheName string, key string, value interface{}) error {
	cachedData, err := c.GetCacheTTL(memCacheName, key)
	if err != nil {
		log.Printf("Error GetCacheTTLJSON for MemCache %s, Key %s, Error : {%+v}\n", memCacheName, key, err)
		return err
	}

	encoded, _ := cachedData.([]byte)

	if err := json.Unmarshal(encoded, value); err != nil {
		return err
	}

	log.Printf("Success GetCacheTTLJSON for MemCache %s, Key %s\n", memCacheName, key)
	return nil
}

func (c *KodingCache) DeleteCache(memCacheName string, key string) error {
	if err := c.MapMemoryTTL[memCacheName].Delete(key); err != nil {
		log.Printf("Error DeleteCache for MemCache %s, Key %s, Error: {%+v}\n", memCacheName, key, err)
		return err
	}

	log.Printf("Success DeleteCache for MemCache %s, Key %s\n", memCacheName, key)
	return nil
}
