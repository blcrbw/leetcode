package kata

func Merge(a <-chan string, b <-chan string) <-chan string {
	c := make(chan string)

	go func() {
		defer close(c)
		for a != nil || b != nil {
			select {
			case m, ok := <-a:
				if ok {
					c <- m
				} else {
					a = nil
				}
			case m, ok := <-b:
				if ok {
					c <- m
				} else {
					b = nil
				}
			}
		}
	}()

	return c
}
