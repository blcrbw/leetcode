package main

// 1115. Print FooBar Alternately
// Suppose you are given the following code:
//      class FooBar {
//        public void foo() {
//          for (int i = 0; i < n; i++) {
//            print("foo");
//          }
//        }
//
//        public void bar() {
//          for (int i = 0; i < n; i++) {
//            print("bar");
//          }
//        }
//      }
//
// The same instance of FooBar will be passed to two different threads:
//
//    thread A will call foo(), while
//    thread B will call bar().
//
// Modify the given program to output "foobar" n times.

type FooBar struct {
	n  int
	ch chan interface{}
}

func NewFooBar(n int) *FooBar {
	return &FooBar{
		n:  n,
		ch: make(chan interface{}),
	}
}

func (fb *FooBar) Foo(printFoo func()) {
	for i := 0; i < fb.n; i++ {
		// printFoo() outputs "foo". Do not change or remove this line.
		printFoo()
		fb.ch <- nil
		<-fb.ch
	}
}

func (fb *FooBar) Bar(printBar func()) {
	for i := 0; i < fb.n; i++ {
		// printBar() outputs "bar". Do not change or remove this line.
		<-fb.ch
		printBar()
		fb.ch <- nil
	}
}
