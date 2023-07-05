package main

import "fmt"

func SumOfSquares(c, quit chan int) {
	n := 1
	for {
		select {
		case c <- n * n:
			n++
		case <-quit:
			return
		}
	}

}

func test_concurrent_sum() {

	mychannel := make(chan int)

	quitchannel := make(chan int)

	sum := 0

	go func() {
		for i := 0; i < 6; i++ {
			sum += <-mychannel
		}
		quitchannel <- 1
		fmt.Println(sum)
	}()
	SumOfSquares(mychannel, quitchannel)
}
