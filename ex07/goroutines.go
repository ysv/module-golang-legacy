package goroutines

func Process(input chan string) chan string {
	var output chan string = make(chan string)
	var done chan bool = make(chan bool)

	go func() {
		str := "(" + <-input + ")"
		output <- string(str)
		done <- true
	}()

	go func() {
		<-done
		close(output)
	}()

	return output
}
