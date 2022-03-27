package engine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type mPrinter struct {
	wg *sync.WaitGroup
	evenCh chan int
	oddCh  chan int
}

func (p *mPrinter) mPrintNumber() {
	defer p.wg.Wait()
	p.wg.Add(2)
	go p.printEven()
	go p.printOdd()
	for i := 1; i <= 10; i ++ {
		if i & 1 == 0 {
			p.evenCh <- i
			time.Sleep(100 * time.Millisecond)
		} else {
			p.oddCh <- i
			time.Sleep(200 * time.Millisecond)
		}
	}
	p.evenCh <- -1
	p.oddCh <- -1
}

func (p *mPrinter) printEven() {
	for {
		select {
		case value := <- p.evenCh:
			if value == -1 {
				p.wg.Done()
				return
			}
			fmt.Printf("even value: %v\n", value)
		}
	}
}

func (p *mPrinter) printOdd() {
	for {
		select {
		case value := <- p.oddCh:
			if value == -1 {
				p.wg.Done()
				return
			}
			fmt.Printf("odd value: %v\n", value)
		}
	}
}

func TestPrintNumber(t *testing.T) {
	mp := &mPrinter{
		wg: &sync.WaitGroup{},
		oddCh: make(chan int, 0),
		evenCh: make(chan int, 0),
	}
	mp.mPrintNumber()
}
