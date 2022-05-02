package once

import (
	"runtime"
	"testing"
)

func Test_Once(t *testing.T) {
	runtime.GOMAXPROCS(runtime.NumCPU())

	totalCount := 10000
	counter := Counter{runCount: 0, initCount: 0}
	done := make(chan struct{})

	for i := 0; i < totalCount; i++ {
		go func(index int) {
			counter.Increment()
			done <- struct{}{}
		}(i)
	}

	for i := 0; i < totalCount; i++ {
		<-done
	}

	if counter.runCount != totalCount {
		t.Error("Invalid runCount")
	}

	if counter.initCount != 1 {
		t.Error("Invalid initCount")
	}
}
