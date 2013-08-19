package main

import (
//"encoding/json"
)

type CacheObj struct {
	Hosts   []string            `json:"hosts"`
	Cluster map[string][]string `json:"cluster"`
	Metrics map[string][]string `json:"metrics"`
}
