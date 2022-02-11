package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	numbers := make([]int, 3)
	numbers[0] = 1
	numbers[1] = 2
	numbers[2] = 3
	var wg sync.WaitGroup
	wg.Add(2)
	go printNumbers(numbers, &wg)
	go printNumbersWithUnderscores(&numbers, &wg)
	numbers = append(numbers, 4)
	wg.Wait()
}

func printNumbers(numbers []int, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(5 * time.Second)
	for _, number := range numbers {
		fmt.Println(number)
	}
}

func printNumbersWithUnderscores(numbers *[]int, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(5 * time.Second)
	for _, number := range *numbers {
		fmt.Printf("__%d\n", number)
	}
}
