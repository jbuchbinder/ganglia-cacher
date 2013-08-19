package main

import (
	"encoding/xml"
	"fmt"
	//"io/ioutil"
	"net"
)

func GetGmetadXml(server string, port int) (GangliaXml, error) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", server, port))
	if err != nil {
		return GangliaXml{}, err
	}

	// Download all content
	/*
		body, err := ioutil.ReadAll(conn)
		if err != nil {
			return GangliaXml{}, err
		}
	*/

	// Parse and unmarshal
	var x GangliaXml
	decoder := xml.NewDecoder(conn)
	decoder.CharsetReader = CharsetReader
	err = decoder.Decode(&x)
	if err != nil {
		return GangliaXml{}, err
	}

	return x, nil
}
