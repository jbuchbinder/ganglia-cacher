package main

import (
	"sort"
)

func ConvertXmlToCache(x GangliaXml) (CacheObj, error) {
	c := CacheObj{}

	hosts := make(map[string]string)

	c.Cluster = make(map[string][]string)
	c.Metrics = make(map[string][]string)

	// Create cluster map
	bCM := NewTimerWithLabel("Cluster map")
	for _, gv := range x.Grid {
		for _, cv := range gv.Cluster {
			for _, hv := range cv.Hosts {
				// Form hosts map
				hosts[hv.Name] = hv.Name

				// Ensure that map exists
				if _, ok := c.Cluster[hv.Name]; ok {
					c.Cluster[hv.Name] = append(c.Cluster[hv.Name], cv.Name)
				} else {
					c.Cluster[hv.Name] = []string{cv.Name}
				}
				// Create metric map
				for _, mv := range hv.Metrics {
					if _, ok := c.Metrics[mv.Name]; ok {
						c.Metrics[mv.Name] = append(c.Metrics[mv.Name], hv.Name)
					} else {
						c.Metrics[mv.Name] = []string{hv.Name}
					}
				}
			}
		}
	}
	bCM.StopLog()

	// Collapse down map values to array
	bMA := NewTimerWithLabel("Map to Array")
	c.Hosts = []string{}
	for _, chv := range hosts {
		if len(chv) > 2 {
			c.Hosts = append(c.Hosts, chv)
		}
	}
	sort.Strings(c.Hosts)
	bMA.StopLog()

	return c, nil
}
