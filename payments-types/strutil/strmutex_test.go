package strutil

import (
	"math/rand"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_StrMutex(t *testing.T) {
	strMutex := NewStrMutex()
	assert.Panics(t, func() { strMutex.Unlock("test") }, "not locked string should panic")

	numParallel := 100
	numAccess := 1000
	wg := sync.WaitGroup{}
	wg.Add(numParallel)

	testFunc := func() {
		for i := 0; i < numAccess; i++ {
			strMutex.Lock("test")
			time.Sleep(time.Nanosecond * time.Duration(rand.Intn(100)))
			strMutex.Unlock("test")
			time.Sleep(1)
		}
		wg.Done()
	}

	for i := 0; i < numParallel; i++ {
		go testFunc()
	}
	wg.Wait()
}
