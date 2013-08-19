package main

import (
	php "github.com/reiver/go-php"
	"io/ioutil"
)

func init() {
	CacheMap["phpfile"] = func() CacheInterface {
		return new(CachePhpFile)
	}
}

type CachePhpFile struct {
	LocalFilename string
}

func (self *CachePhpFile) Connect() {
}

func (self *CachePhpFile) Write(obj CacheObj) error {
	b := php.Serialize(obj)
	err := ioutil.WriteFile(self.LocalFilename, []byte(b), 0666)
	if err != nil {
		return err
	}
	return nil
}
