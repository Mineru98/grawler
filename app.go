package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"

	"./target/naver"
)

func main() {
	cpuNumber := 2
	runtime.GOMAXPROCS(cpuNumber)
	var wait sync.WaitGroup
	wait.Add(cpuNumber)

	startTime := time.Now()

	go func() {
		for i := 0; i < 15; i++ {
			naver.Crawler(i)
		}

		defer wait.Done()
	} ()

	go func() {
		for i := 0; i < 15; i++ {
			naver.Crawler(i)
		}

		defer wait.Done()
	} ()

	wait.Wait()
	elapsedTime := time.Since(startTime)
	
	fmt.Printf("실행시간: %s\n", elapsedTime)
}