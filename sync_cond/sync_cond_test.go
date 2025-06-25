package synccond

import (
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestBuffer(t *testing.T) {
	buffer := NewBuffer()

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	wg := &sync.WaitGroup{}

	// 生产者
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 50; i++ {
			if err := buffer.Producer(r.Intn(100), i); err != nil {
				t.Error(err)
			}
			time.Sleep(time.Millisecond * 100)
		}

	}()

	// 消费者
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 50; i++ {
			buffer.Consumer(i)
			time.Sleep(time.Millisecond * 300)
		}
	}()

	wg.Wait()

}
