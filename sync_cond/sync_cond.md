### sync.Cond

sync.Cond 这个包主要用于实现不同协程之间的同步

在创建一个cond变量的时候需要在构造函数里面传入一个Locker接口, 通常使用sync.Mutex当作传入的Locker

例如: cond := sync.NewCond(&sync.Mutex)

    Go 标准库中 sync.Cond 的核心字段：
    type Cond struct {
        L Locker       // 关联的锁
        notify  notifyList  // 等待队列
        ... // 其他未导出的字段
    }

#### sync.Wait

当一个 goroutine 调用 cond.Wait() 时，它会自动执行以下操作：

1. 释放当前持有的锁

2. 阻塞自己，进入等待状态，直到被其他 goroutine 调用 Signal() 或 Broadcast() 唤醒

3. 被唤醒后，Wait() 会重新尝试获取锁，成功获得锁后，Wait() 返回，goroutine 继续执行


#### sync.Signal

1. Signal() 的作用是唤醒一个正在等待该条件变量的 goroutine。如果没有 goroutine 在等待，调用 Signal() 不会有任何效果

2. Signal() 不释放锁，调用者仍然持有锁直到显式调用 Unlock()

3. 被唤醒的 goroutine 在 Wait() 返回之前，必须重新获得锁

4. 如果多个 goroutine 在等待，Signal() 只唤醒一个


#### sync.Broadcast

1. 调用 Broadcast() 时，调用者必须持有锁

2. Broadcast() 不释放锁，调用者仍然持有锁直到显式调用 Unlock()

3. 被唤醒的所有 goroutine 都会从 Wait() 返回，之后开始竞争锁







