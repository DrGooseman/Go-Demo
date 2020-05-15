package main

import (
	"fmt"
	"time"
	"sync"
)

func main() {

	//worker pools

	jobs := make(chan int, 44)
	results := make(chan int, 44)

	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)

	for i := 0; i < 44; i++{
		jobs <- i
	}
	close(jobs)

	for j := 0; j < 44; j++{
		fmt.Println(<-results)
	}

		// channels

		c := make(chan string)
		go count2("sheep", c)
		
		for msg := range c {
			fmt.Println(msg)
		}

// 2 channels

c1 := make(chan string)
c2 := make(chan string)

go func() {
	for {
		c1 <- "Every 500ms"
		time.Sleep(time.Millisecond * 500)
	}
} ()

go func() {
	for {
		c2 <- "Every 2 seconds"
		time.Sleep(time.Second * 2)
	}
} ()

for {
	select {
	case msg1 := <- c1:
		fmt.Println(msg1)
	case msg2 := <- c2:
		fmt.Println(msg2)
	}
}




		// WaitGroup

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		count("sheep")
		wg.Done()
	}()

	go func() {
		count("fish")
		wg.Done()
	}()

	wg.Wait()
}

func worker(jobs <- chan int, results chan <- int){
	for n := range jobs {
		results <- fib(n)
	}
}

func fib (n int) int {
	if n <= 1 {
		return n
	}

	return fib(n - 1) + fib(n - 2)
}


func count(thing string) {
	for i := 1; i <= 5; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}

func count2(thing string, c chan string) {
	for i := 1; i <= 5; i++ {
		c <- thing
		time.Sleep(time.Millisecond * 500)
	}

	close(c)
}