package main

import (
	"testing"
	"time"
)

func Test_dine(t *testing.T) {
	eatTime = 0 * time.Second   // время для трапизы
	thinkTime = 0 * time.Second // время на подумать
	sleepTime = 0 * time.Second // время для сна

	for i := 0; i < 10; i++ {
		OrderFinished = []string{}
		dine()
		if len(OrderFinished) != 5 {
			t.Errorf("incorrect length of slice; expected 5 but got %d", len(OrderFinished))
		}
	}
}

func Test_dineWithVaryingDelays(t *testing.T) {
	var theTests = []struct {
		name  string
		delay time.Duration
	}{
		{"zero delay", time.Second * 0},
		{"quarter second delay", time.Microsecond * 250},
		{"half second delay", time.Microsecond * 500},
	}

	for _, e := range theTests {
		OrderFinished = []string{}
		eatTime = e.delay
		thinkTime = e.delay
		sleepTime = e.delay
		dine()
		if len(OrderFinished) != 5 {
			t.Errorf("%s, incorrect length of slice; expected 5 but got %d", e.name, len(OrderFinished))
		}
	}
}
