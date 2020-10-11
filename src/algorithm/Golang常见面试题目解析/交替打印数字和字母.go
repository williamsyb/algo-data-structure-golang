package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	numberCh, letterCh := make(chan int), make(chan int)
	go func() {
		i := 1
		for {
			time.Sleep(time.Duration(300) * time.Millisecond)
			select {
			case <-numberCh:
				fmt.Print(i)
				i++
				fmt.Print(i)
				i++
				letterCh <- 1
			}
		}
	}()

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		A, Z := 'A', 'Z'
		ANum, ZNum := int(A), int(Z)
		i := ANum
		for {
			time.Sleep(time.Duration(300) * time.Millisecond)
			select {
			case <-letterCh:
				if i > ZNum {
					wg.Done()
					return
				}
				fmt.Print(string(rune(i)))
				i++
				fmt.Print(string(rune(i)))
				i++
				numberCh <- 1
			}
		}
	}()

	numberCh <- 1
	wg.Wait()
}
