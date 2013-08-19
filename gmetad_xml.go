package main

import (
	"encoding/xml"
	"math"
)

type GangliaXml struct {
	XMLName xml.Name  `xml:"GANGLIA_XML"`
	Version string    `xml:"VERSION,attr"`
	Source  string    `xml:"SOURCE,attr"`
	Grid    []GridObj `xml:"GRID"`
}

type GridObj struct {
	Name      string       `xml:"NAME,attr"`
	Authority string       `xml:"AUTHORITY,attr"`
	Localtime int64        `xml:"LOCALTIME,attr"`
	Cluster   []ClusterObj `xml:"CLUSTER"`
}

type ClusterObj struct {
	Name      string    `xml:"NAME,attr"`
	Localtime float64   `xml:"LOCALTIME,attr"`
	Owner     string    `xml:"OWNER,attr"`
	LatLong   string    `xml:"LATLONG,attr"`
	Url       string    `xml:"URL,attr"`
	Hosts     []HostObj `xml:"HOST"`
}

type HostObj struct {
	Name     string      `xml:"NAME,attr"`
	IP       string      `xml:"IP,attr"`
	Location string      `xml:"LOCATION,attr"`
	Tn       int         `xml:"TN,attr"`
	Tmax     int         `xml:"TMAX,attr"`
	Dmax     int         `xml:"DMAX,attr"`
	Reported float64     `xml:"REPORTED,attr"`
	Metrics  []MetricObj `xml:"METRIC"`
}

func (self *HostObj) HostAlive(cluster ClusterObj) bool {
	ttl := 60
	if self.Tn > 0 && self.Tmax > 0 {
		if self.Tn > self.Tmax*4 {
			return false
		}
	} else {
		if math.Abs(cluster.Localtime-self.Reported) > float64(ttl*4) {
			return false
		}
	}
	return true
}

type MetricObj struct {
	Name      string         `xml:"NAME,attr"`
	Value     string         `xml:"VAL,attr"`
	Type      string         `xml:"TYPE,attr"`
	Units     string         `xml:"UNITS,attr"`
	Tn        int            `xml:"TN,attr"`
	Tmax      int            `xml:"TMAX,attr"`
	Dmax      int            `xml:"DMAX,attr"`
	Slope     string         `xml:"SLOPE,attr"`
	ExtraData []ExtraDataObj `xml:"EXTRA_DATA>EXTRA_ELEMENT"`
}

type ExtraDataObj struct {
	Name  string `xml:"NAME,attr"`
	Value string `xml:"VAL,attr"`
}
