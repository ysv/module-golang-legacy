//package goroutines
package main

import (
	"fmt"
	//"time"
	"sync"
)

var wg sync.WaitGroup

func Process(input chan string) chan string {
	s := <-input
	var output chan string = make(chan string)
	wg.Add(1)
	go in(s, output)
	wg.Wait()
	go out(input)

	return output
}

func in(s string, output chan string) {
	defer wg.Done()
	str := make([]byte, len(s)+2)
	str[0] = 40
	for i := 0; i < len(s); i++ {
		fmt.Println(string(str))
		str[i+1] = s[i]
		if i+2 > len(s) {
			str[i+2] = 41
		}
	}
	fmt.Println(string(str))
	output <- string(str)
}

func out(input chan string) {
	close(input)
}

func main() {
	var input chan string = make(chan string)
	input <- "hello world!"
	Process(input)
}
