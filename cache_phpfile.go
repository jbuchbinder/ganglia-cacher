package main

import (
	php "github.com/reiver/go-php"
	"io/ioutil"
)

// TODO: FIXME: PHP serialization is busted, since it doesn't accept reflection for objects

func init() {
	CacheMap["phpfile"] = func() CacheInterface {
		return new(CachePhpFile)
	}
}

type CachePhpFile struct {
	LocalFilename string
}

func (self *CachePhpFile) Configure(conf string) {
	self.LocalFilename = conf
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
