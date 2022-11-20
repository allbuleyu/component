package goroutine

import (
	"sync"
	"sync/atomic"
)

type Multi struct {
	workChan chan func() error
	done     chan struct{}

	wg     *WaitGroup
	errNum int32
	err    error
	once   *sync.Once
}

func NewMulti(taskNum int) *Multi {
	if taskNum == 0 {
		taskNum = 3
	}

	multi := &Multi{
		workChan: make(chan func() error, taskNum*3),
		done:     make(chan struct{}, 0),
		wg:       &WaitGroup{},
		err:      nil,
		errNum:   0,
		once:     &sync.Once{},
	}

	for i := 0; i < taskNum; i++ {
		multi.wg.Warp(multi.receiveWork)
	}

	return multi
}

func (m *Multi) Run(f func() error) {
	m.sendWork(f)
}

func (m *Multi) receiveWork() {
	for {
		select {
		case f := <-m.workChan:
			if err := f(); err != nil {
				atomic.AddInt32(&m.errNum, 1)
				m.once.Do(func() {
					m.err = err
				})
			}
		case <-m.done:
			break
		}
	}
}

func (m *Multi) sendWork(f func() error) {
	select {
	case m.workChan <- f:
	case <-m.done:
	}
}

func (m *Multi) Close() {
	m.close()
}

func (m *Multi) close() {
	close(m.done)
}

func (m *Multi) Wait() error {
	m.wg.Wait()
	close(m.workChan)

	return m.err
}
