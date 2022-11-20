package goroutine

import (
	"log"
	"runtime/debug"
	"sync"
)

func Safe(f func()) func() {
	return func() {
		defer func() {
			if err := recover(); err != nil {
				stack := debug.Stack()
				log.Printf("go routine panic: %v, %s", err, stack)
			}
		}()

		f()
	}
}

type WaitGroup struct {
	sync.WaitGroup
}

func (wg *WaitGroup) Warp(f func()) {
	wg.Add(1)

	go Safe(func() {
		defer wg.Done()
		f()
	})()
}
