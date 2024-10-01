package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func foo() {
	for i := 0; i <= 10; i++ {
		fmt.Println("FOO")
	}
	wg.Done()
}

func bar() {
	for j := 0; j <= 10; j++ {
		fmt.Println("BAR")
	}
}

func main() {
	fmt.Println("OS\t", runtime.GOOS)
	fmt.Println("ARCH\t", runtime.GOARCH)
	fmt.Println("CPUs\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	wg.Add(1)
	go foo()
	bar()
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	wg.Wait()

}
