package main

import (
	"fmt"
	"time"
)

func main() {
	s := NewChanTest()

	i := 0

	go func() {
		s.update1()
	}()
	go func() {
		s.update2()
	}()
	for i < 1000 {
		s.sendValue(i)
		i++
	}
	time.Sleep(10 * time.Second)
}

type ChanTest struct {
	Buf1 chan int
	Buf2 chan int

	Ticker1 *time.Ticker
	Ticker2 *time.Ticker
}

func NewChanTest() *ChanTest {
	return &ChanTest{
		Buf1:    make(chan int, 400),
		Buf2:    make(chan int, 400),
		Ticker1: time.NewTicker(time.Second),
		Ticker2: time.NewTicker(time.Second),
	}
}

func (s *ChanTest) close() {
	s.Ticker1.Stop()
	s.Ticker2.Stop()
	close(s.Buf1)
	close(s.Buf2)
}

func (s *ChanTest) update1() {
	for {
		select {
		case <-s.Ticker1.C:
			fmt.Println("ticker 1", s.getValue(s.Buf1))
		default:
		}
	}
}

func (s *ChanTest) update2() {
	for {
		select {
		case <-s.Ticker2.C:
			fmt.Println("ticker 2", s.getValue(s.Buf2))
		default:
		}
	}
}

func (s *ChanTest) getValue(ch chan int) []int {
	var ans []int
	l := len(ch)
	for i := 0; i < l; i++ {
		v, ok := <-ch
		if !ok {
			return ans
		}
		ans = append(ans, v)
	}
	return ans
}

func (s *ChanTest) sendValue(value int) {
	if value%2 != 0 {
		s.Buf1 <- value
		return
	} else {
		s.Buf2 <- value
		return
	}
	fmt.Println("buffer is full")
}
