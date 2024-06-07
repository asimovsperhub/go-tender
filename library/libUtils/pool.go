package libUtils

import "sync"

// pool

type WaitGroup struct {
	workChan chan int
	// WaitGroup等待一组goroutines完成。
	//主goroutine调用Add来设置
	// goroutines等待。然后是每一个goroutine
	//在完成时运行并调用Done。与此同时，
	//等待可以用来阻塞，直到所有goroutines完成
	wg sync.WaitGroup
}

func NewPool(coreNum int) *WaitGroup {
	ch := make(chan int, coreNum)
	return &WaitGroup{
		workChan: ch,
		wg:       sync.WaitGroup{},
	}
}
func (ap *WaitGroup) Add(num int) {
	for i := 0; i < num; i++ {
		ap.workChan <- i
		// Add将增量添加到WaitGroup计数器，增量可以是负数。
		// 如果计数器变为0，所有在Wait阻塞的goroutines将被释放。
		ap.wg.Add(1)
	}
}
func (ap *WaitGroup) Done() {
LOOP:
	for {
		select {
		case <-ap.workChan:
			break LOOP
		}
	}
	ap.wg.Done()
}

func (ap *WaitGroup) Wait() {
	ap.wg.Wait()
}
