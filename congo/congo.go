package congo

import (
	"fmt"
	"sync"
)

func add(a, b int, wg *sync.WaitGroup, ch chan string) {
	result := a + b
	ch <- fmt.Sprintf("%d + %d = %d", a, b, result)
	wg.Done()
}

func subtract(a, b int, wg *sync.WaitGroup, ch chan string) {
	result := a - b
	ch <- fmt.Sprintf("%d - %d = %d", a, b, result)
	wg.Done()
}

func multiply(a, b int, wg *sync.WaitGroup, ch chan string) {
	result := a * b
	ch <- fmt.Sprintf("%d * %d = %d", a, b, result)
	wg.Done()
}

func divide(a, b int, wg *sync.WaitGroup, ch chan string) {
	if b != 0 {
		result := a / b
		ch <- fmt.Sprintf("%d / %d = %d", a, b, result)
	} else {
		ch <- "Division by zero!"
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan string)

	a, b := 10, 5

	wg.Add(4) // we have functions here

	go add(a, b, &wg, ch)
	go subtract(a, b, &wg, ch)
	go multiply(a, b, &wg, ch)
	go divide(a, b, &wg, ch)

	go func() {
		wg.Wait()
		close(ch)
	}()

	for result := range ch {
		fmt.Println(result)
	}
}
