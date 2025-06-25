package synccond

import (
	"fmt"
	"sync"
)

type Buffer struct {
	mtx   *sync.Mutex
	cond  *sync.Cond
	data  []int
	limit int
}

func NewBuffer() *Buffer {
	mtx := &sync.Mutex{}
	return &Buffer{mtx, sync.NewCond(mtx), make([]int, 0), 10}
}

func (b *Buffer) Producer(value, i int) error {
	b.mtx.Lock()
	defer b.mtx.Unlock()

	// 尝试往缓冲区生产物品
	if len(b.data) == b.limit {
		// 缓冲区已满, 挂起阻塞
		b.cond.Wait()
		fmt.Println("No.", i, "producer was waited")
	}

	// 缓冲区未满, 往缓冲区生产
	b.data = append(b.data, value)

	// Signal信号, 唤醒cond等待队列中的一个协程
	b.cond.Signal()
	fmt.Println("No.", i, "producer wake up")

	return nil
}

func (b *Buffer) Consumer(i int) int {
	b.mtx.Lock()
	defer b.mtx.Unlock()

	// 尝试从缓冲区中消费物品
	if len(b.data) == 0 {
		// 缓冲区为空, 挂起阻塞
		b.cond.Wait()
		fmt.Println("No.", i, "consumer was waited")
	}

	// 消费一件物品
	res := b.data[0]
	b.data = b.data[1:]

	// Signal信号, 唤醒cond等待队列中的一个协程
	b.cond.Signal()
	fmt.Println("No.", i, "consumer wake up")

	return res
}
