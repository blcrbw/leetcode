package main

import (
	"fmt"
	"sync"
)

// 1116. Print Zero Even Odd
//
// You have a function printNumber that can be called with an integer parameter and prints it to the console.
//
//    For example, calling printNumber(7) prints 7 to the console.
//
// You are given an instance of the class ZeroEvenOdd that has three functions: zero, even, and odd. The same instance
// of ZeroEvenOdd will be passed to three different threads:
//
//    Thread A: calls zero() that should only output 0's.
//    Thread B: calls even() that should only output even numbers.
//    Thread C: calls odd() that should only output odd numbers.
//
// Modify the given class to output the series "010203040506..." where the length of the series must be 2n.
//
// Implement the ZeroEvenOdd class:
//
//    ZeroEvenOdd(int n) Initializes the object with the number n that represents the numbers that should be printed.
//    void zero(printNumber) Calls printNumber to output one zero.
//    void even(printNumber) Calls printNumber to output one even number.
//    void odd(printNumber) Calls printNumber to output one odd number.

type ZeroEvenOdd struct {
	zc    chan *int
	ec    chan *int
	oc    chan *int
	pc    chan interface{}
	mutex *sync.Mutex
}

func NewZeroEvenOdd() *ZeroEvenOdd {
	zc := make(chan *int)
	ec := make(chan *int)
	oc := make(chan *int)
	pc := make(chan interface{})
	return &ZeroEvenOdd{
		zc,
		ec,
		oc,
		pc,
		&sync.Mutex{},
	}
}

func (o *ZeroEvenOdd) zero() {
	for i := range o.zc {
		fmt.Print("0")
		if *i%2 == 0 {
			o.oc <- i
		} else {
			o.ec <- i
		}
	}
}

func (o *ZeroEvenOdd) even() {
	for i := range o.ec {
		fmt.Print(*i)
		o.pc <- nil
	}
}

func (o *ZeroEvenOdd) odd() {
	for i := range o.oc {
		fmt.Print(*i)
		o.pc <- nil
	}
}

func (o *ZeroEvenOdd) init() {
	go func() { o.zero() }()
	go func() { o.even() }()
	go func() { o.odd() }()
}

func (o *ZeroEvenOdd) printNumber(n int) {
	o.mutex.Lock()
	for i := 1; i <= n; i++ {
		o.zc <- &i
		<-o.pc
	}
	o.mutex.Unlock()
}

func main() {
	zeo := NewZeroEvenOdd()
	zeo.init()

	fmt.Print("printNumber(1): ")
	zeo.printNumber(1)
	fmt.Print("\nprintNumber(7): ")
	zeo.printNumber(7)
	fmt.Print("\nprintNumber(13): ")
	zeo.printNumber(13)
}
