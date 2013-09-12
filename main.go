package main

import (
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"log"
)

var (
	Debug              = flag.Bool("debug", false, "Debug printing")
	Server             = flag.String("server", "127.0.0.1", "Gmetad server")
	Port               = flag.Int("port", 8651, "Gmetad server port")
	TestMode           = flag.Bool("test", false, "Test run mode")
	Cache              = flag.String("cache", "", "Cache module")
	CacheConfiguration = flag.String("cacheconf", "", "Cache module configuration")
	ListCacheModules   = flag.Bool("listcache", false, "List cache modules")
)

func main() {
	flag.Parse()

	if *ListCacheModules {
		for k, _ := range CacheMap {
			fmt.Println(k)
		}
		return
	}

	if *TestMode {
		// Single run
		gx, err := GetGmetadXml(*Server, *Port)
		if err != nil {
			log.Fatalln(err)
		}
		if *Debug {
			pt, _ := xml.MarshalIndent(gx, " ", "  ")
			fmt.Println(string(pt))
		}

		if *Debug {
			fmt.Println("##### CONVERTING TO JSON CACHE #####")
		}

		c, err := ConvertXmlToCache(gx)
		if err != nil {
			log.Fatalln(err)
		}

		if *Debug {
			jt, _ := json.MarshalIndent(c, " ", "  ")
			fmt.Println(string(jt))
		}

		return
	}

	// Check command line options
	if *Cache == "" {
		log.Panicln("No cache option specified")
	}

	// Standard execution starts here
	log.Print("Beginning cache run")

	log.Printf("Connecting to %s:%d", *Server, *Port)
	b := NewTimer()
	gx, err := GetGmetadXml(*Server, *Port)
	b.StopLog()
	if err != nil {
		log.Panicln(err)
	}

	log.Print("Converting to cache format")
	b = NewTimerWithLabel("ConvertXmlToCache")
	c, err := ConvertXmlToCache(gx)
	b.StopLog()
	if err != nil {
		log.Panicln(err)
	}
	log.Print("Completed converting to cache format")

	log.Printf("Using cache module %s", *Cache)
	cache := GetCache(*Cache)
	cache.Configure(*CacheConfiguration)
	log.Print("Calling Connect()")
	cache.Connect()
	log.Print("Calling Write()")
	bW := NewTimerWithLabel("Write()")
	cache.Write(c)
	bW.StopLog()
	log.Print("Completed cache run")
}
