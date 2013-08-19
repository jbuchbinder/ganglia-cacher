package main

import (
	"fmt"
	"log"
	"strings"
)

const (
	MEMCACHE_CACHE_PREFIX     = "ganglia_cache_"
	MEMCACHE_TIMESTAMP_PREFIX = "ganglia_cache_timestamp_"
)

var (
	CacheMap = map[string]func() CacheInterface{}
)

type CacheInterface interface {
	Configure(string)
	Connect()
	Write(CacheObj) error
}

// Resolves CacheInterface objects based on their string names.
func GetCache(c string) CacheInterface {
	log.Print("Selecting cache backend using: '" + c + "'")
	cName := strings.TrimSpace(c)
	if _, exists := CacheMap[cName]; exists {
		return CacheMap[cName]()
	} else {
		fmt.Println("Unable to resolve cache backend " + cName)
	}
	return nil
}
