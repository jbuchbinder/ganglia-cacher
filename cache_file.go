package main

import (
	"encoding/json"
	"io/ioutil"
)

func init() {
	CacheMap["file"] = func() CacheInterface {
		return new(CacheFile)
	}
}

type CacheFile struct {
	LocalFilename string
}

func (self *CacheFile) Connect() {
}

func (self *CacheFile) Write(obj CacheObj) error {
	b, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(self.LocalFilename, b, 0666)
	if err != nil {
		return err
	}
	return nil
}
