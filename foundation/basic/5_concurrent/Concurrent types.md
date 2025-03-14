涵盖的并发方式:

1.goroutine（基础并发）
2.channel（消息通信）
3.sync.WaitGroup（等待多个 goroutine 完成）
4.sync.Mutex & sync.RWMutex（互斥锁，读写锁）
5.sync.Cond（条件变量）
6.sync.Once（只执行一次）
7.sync.Pool（对象复用池）
8.atomic 操作（原子操作，锁替代方案）
9.context 控制并发（超时、取消）