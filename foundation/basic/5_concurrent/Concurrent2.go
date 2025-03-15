package main

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

// 1. Goroutine 示例
func GoroutineExample() {
	go func() {
		fmt.Println("Hello from Goroutine!")
	}()
	time.Sleep(100 * time.Millisecond) // 等待 goroutine 执行
}

// 2. Channel 示例
func ChannelExample() string {
	ch := make(chan string)
	go func() {
		ch <- "Hello from Channel!"
	}()
	return <-ch
}

// 3. sync.WaitGroup 示例
func WaitGroupExample() int {
	var wg sync.WaitGroup
	sum := 0
	mu := sync.Mutex{}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			mu.Lock()
			sum += i
			mu.Unlock()
		}(i)
	}

	wg.Wait()
	return sum
}

// 4. sync.Mutex & sync.RWMutex 示例
type SafeCounter struct {
	mu    sync.RWMutex
	value int
}

func (c *SafeCounter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *SafeCounter) Get() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.value
}

// 5. sync.Cond 示例
func CondExample() string {
	cond := sync.NewCond(&sync.Mutex{})
	data := ""
	done := false

	go func() {
		time.Sleep(100 * time.Millisecond)
		cond.L.Lock()
		data = "Data is Ready"
		done = true
		cond.Signal()
		cond.L.Unlock()
	}()

	cond.L.Lock()
	for !done {
		cond.Wait()
	}
	cond.L.Unlock()

	return data
}

// 6. sync.Once 示例
func OnceExample() string {
	var once sync.Once
	var result string

	onceFunc := func() {
		result = "Initialized"
	}

	once.Do(onceFunc)
	return result
}

// 7. sync.Pool 示例
func PoolExample() string {
	pool := sync.Pool{
		New: func() interface{} {
			return "New Object"
		},
	}

	obj1 := pool.Get().(string)
	pool.Put("Reused Object")
	obj2 := pool.Get().(string)

	return obj1 + " | " + obj2
}

// 8. atomic 操作 示例
func AtomicExample() int32 {
	var counter int32
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt32(&counter, 1)
		}()
	}

	wg.Wait()
	return counter
}

// 9. context 控制并发（超时示例）
func ContextTimeoutExample() string {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	ch := make(chan string)
	go func() {
		time.Sleep(200 * time.Millisecond)
		ch <- "Finished Work"
	}()

	select {
	case msg := <-ch:
		return msg
	case <-ctx.Done():
		return "Timeout"
	}
}

// 9. context 控制并发（取消示例）
func ContextCancelExample() string {
	ctx, cancel := context.WithCancel(context.Background())
	ch := make(chan string)

	go func() {
		time.Sleep(100 * time.Millisecond)
		select {
		case <-ctx.Done():
			ch <- "Cancelled"
		default:
			ch <- "Finished Work"
		}
	}()

	cancel()
	return <-ch
}

func TestChannelExample(t *testing.T) {
	got := ChannelExample()
	want := "Hello from Channel!"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestWaitGroupExample(t *testing.T) {
	got := WaitGroupExample()
	want := 10 // 0+1+2+3+4 = 10
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestSafeCounter(t *testing.T) {
	counter := SafeCounter{}
	counter.Inc()
	counter.Inc()
	got := counter.Get()
	want := 2
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestCondExample(t *testing.T) {
	got := CondExample()
	want := "Data is Ready"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestOnceExample(t *testing.T) {
	got := OnceExample()
	want := "Initialized"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestPoolExample(t *testing.T) {
	got := PoolExample()
	want := "New Object | Reused Object"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestAtomicExample(t *testing.T) {
	got := AtomicExample()
	var want int32 = 5
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestContextTimeoutExample(t *testing.T) {
	got := ContextTimeoutExample()
	want := "Timeout"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestContextCancelExample(t *testing.T) {
	got := ContextCancelExample()
	want := "Cancelled"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
