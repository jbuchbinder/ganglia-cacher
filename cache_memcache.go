package main

import (
	"encoding/json"
	"fmt"
	"github.com/bradfitz/gomemcache/memcache"
	"os"
	"strings"
	"time"
)

func init() {
	CacheMap["memcache"] = func() CacheInterface {
		return new(CacheMemcache)
	}
}

type CacheMemcache struct {
	Servers []string
	Client  *memcache.Client
}

func (self *CacheMemcache) Configure(conf string) {
	self.Servers = strings.Split(conf, ",")
}

func (self *CacheMemcache) Connect() {
	ss := new(memcache.ServerList)
	ss.SetServers(self.Servers...)
	self.Client = memcache.NewFromSelector(ss)
}

func (self *CacheMemcache) Write(obj CacheObj) error {
	hn, err := os.Hostname()
	if err != nil {
		return err
	}
	b, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	err = self.Client.Set(&memcache.Item{Key: MEMCACHE_CACHE_PREFIX + hn, Value: b})
	if err != nil {
		return err
	}
	err = self.Client.Set(&memcache.Item{Key: MEMCACHE_TIMESTAMP_PREFIX + hn, Value: []byte(fmt.Sprintf("%d", time.Now().Unix()))})
	return err
}
