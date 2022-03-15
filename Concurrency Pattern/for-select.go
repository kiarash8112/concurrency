package main

func main() {
	stringStream := make(chan interface{})
	done := make(chan interface{}, 2)
	for _, s := range []string{"a", "b", "c"} {
		select {
		case <-done:
			return
		case stringStream <- s:
		}
	}
}
