package main

import (
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
)

var (
	Debug    = flag.Bool("debug", false, "Debug printing")
	Server   = flag.String("server", "127.0.0.1", "Gmetad server")
	Port     = flag.Int("port", 8651, "Gmetad server port")
	TestMode = flag.Bool("test", false, "Test run mode")
)

func main() {
	flag.Parse()

	if *TestMode {
		// Single run
		gx, err := GetGmetadXml(*Server, *Port)
		if err != nil {
			panic(err)
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
			panic(err)
		}

		if *Debug {
			jt, _ := json.MarshalIndent(c, " ", "  ")
			fmt.Println(string(jt))
		}

		return
	}

	// Standard execution starts here
}
