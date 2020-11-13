program

package main

import (
	"fmt"
	"sync"
)

func countRoutine(upTo int, wg *sync.WaitGroup) {
	for i := 0; i < upTo; i++ {
		fmt.Println("count routine: ", i)
	}
	wg.Done()
}

func count(upTo int) {
	for i := 0; i < upTo; i++ {
		fmt.Println("count: ", i)
	}
}

func main() {
	var wg sync.WaitGroup

	wg.Add(1)

	go countRoutine(10, &wg)

	count(5)

	wg.Wait()
}
output:
count:  0
count:  1
count:  2
count:  3
count:  4
count routine:  0
count routine:  1
count routine:  2
count routine:  3
count routine:  4
count routine:  5
count routine:  6
count routine:  7
count routine:  8
count routine:  9
