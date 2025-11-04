package kata

func Merge(c ...chan string) <-chan string {
	res := make(chan string)
	done := make(chan struct{})

	num := len(c)

	for _, ch := range c {
		go func(c chan string) {
			for msg := range c {
				res <- msg
			}
			done <- struct{}{}
		}(ch)
	}
	go func() {
		for i := 0; i < num; i++ {
			<-done
		}
		close(done)
		close(res)
	}()

	return res
}
