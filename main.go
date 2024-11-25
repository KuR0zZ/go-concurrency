package main

import (
	"fmt"
	"sync"
)

// Print number 1 to 10
func printNumbers() {
	for i := 0; i < 10; i++ {
		fmt.Println(i + 1)
	}
}

// Print letter a to j
func printLetters() {
	letters := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	for _, letter := range letters {
		fmt.Println(letter)
	}
}

// Send numbers to channel
func produce(c chan int) {
	for i := 0; i < 10; i++ {
		c <- i + 1
	}
	close(c)
}

// Consume numbers from channel
func consume(c chan int) {
	for v := range c {
		fmt.Println(v)
	}
}

func main() {
	// Task 1
	go printNumbers()
	go printLetters()

	// Task 2
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer wg.Done()
		printNumbers()
	}()
	go func() {
		defer wg.Done()
		printLetters()
	}()
	wg.Wait()

	// Task 3
	c := make(chan int)

	go produce(c)
	consume(c)

	// Task 4
	// Buffered channel allow values to be sent based on the set buffered without a corresponding receiver
	// While unbuffered channel must have a corresponding receiver after each sent
	bc := make(chan int, 5)

	go produce(bc)
	consume(bc)

	// Task 5 & 6
	odd := make(chan int)
	even := make(chan int)
	error := make(chan int)

	go func() {
		for i := 1; i < 30; i++ {
			if i > 20 {
				error <- i
				continue
			}
			if i%2 == 0 {
				even <- i
				continue
			}
			odd <- i
		}
		close(even)
		close(odd)
		close(error)
	}()

OuterLoop:
	for {
		select {
		case msg1, ok := <-even:
			if !ok {
				break OuterLoop
			}
			fmt.Println("Received an even number: ", msg1)
		case msg2, ok := <-odd:
			if !ok {
				break OuterLoop
			}
			fmt.Println("Received an odd number: ", msg2)
		case msg3, ok := <-error:
			if !ok {
				break OuterLoop
			}
			fmt.Printf("Error %d is greater than 20\n", msg3)
		}
	}
}
