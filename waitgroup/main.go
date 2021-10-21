package main

import (
	"fmt"
	"sync"
)

func main()  {
	var wg sync.WaitGroup

	fmt.Println("app started")
	wg.Add(2)

	go func1(&wg)
	go func2(&wg)

	wg.Wait()
	fmt.Println("app finished")
}

func func1(wg *sync.WaitGroup)  {
	fmt.Println("function 1")
	wg.Done()
}

func func2(wg *sync.WaitGroup)  {
	fmt.Println("function 2")
	wg.Done()
}
