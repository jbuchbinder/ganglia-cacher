package main

import (
	"log"
	"time"
)

type Benchmark struct {
	Timer time.Time
	Label string
}

func NewTimer() Benchmark {
	b := Benchmark{}
	b.Timer = time.Now()
	return b
}

func NewTimerWithLabel(label string) Benchmark {
	b := Benchmark{}
	b.Timer = time.Now()
	b.Label = label
	return b
}

func (self *Benchmark) StopLog() {
	e := time.Since(self.Timer)
	if self.Label != "" {
		log.Printf("[%s] Duration : %s", self.Label, e.String())
	} else {
		log.Printf("Duration : %s", e.String())
	}
}
