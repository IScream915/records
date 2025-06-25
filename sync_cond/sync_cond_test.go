package synccond

import (
	"math/rand"
	"testing"
	"time"
)

func TestBuffer(t *testing.T) {
	buffer := NewBuffer()

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// 生产者
	go func() {
		for i := 0; i < 50; i++ {
			buffer.Producer(r.Intn(100), i)
			// time.Sleep(time.Millisecond * 500)
		}

	}()

	// 消费者
	go func() {
		for i := 0; i < 50; i++ {
			buffer.Consumer(i)
			// time.Sleep(time.Millisecond * 500)
		}
	}()

	time.Sleep(time.Second * 20)
}
