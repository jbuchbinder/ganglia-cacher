package main

import (
	"encoding/xml"
	"fmt"
	"net"
)

func GetGmetadXml(server string, port int) (GangliaXml, error) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", server, port))
	if err != nil {
		return GangliaXml{}, err
	}

	// Parse and unmarshal
	var x GangliaXml
	decoder := xml.NewDecoder(conn)
	decoder.CharsetReader = CharsetReader
	b := NewTimerWithLabel("Decode XML")
	err = decoder.Decode(&x)
	b.StopLog()
	if err != nil {
		return GangliaXml{}, err
	}

	return x, nil
}
