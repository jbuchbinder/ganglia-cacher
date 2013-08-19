package main

import (
	"log"
	"time"
)

type Benchmark struct {
	Timer time.Time
}

func NewTimer() Benchmark {
	b := Benchmark{}
	b.Timer = time.Now()
	return b
}

func (self *Benchmark) StopLog() {
	e := time.Since(self.Timer)
	log.Printf("Duration : %s", e.String())
}
