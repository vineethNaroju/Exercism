package main

import (
	"fmt"
	"sync"
)

func haha() {
	m := map[rune]int{}

	text := []string{
		"Hello world !",
		"Ahahaha !",
	}

	lock := make(chan bool, 1)

	wg := sync.WaitGroup{}

	for _, str := range text {

		wg.Add(1)

		go func(s string) {
			defer wg.Done()

			for _, x := range s {
				lock <- true
				m[x]++
				<-lock
			}
		}(str)
	}

	wg.Wait()

	fmt.Print(m)
}

func GetFreq(s string) map[rune]int {
	m := map[rune]int{}

	for _, r := range s {
		m[r]++
	}

	return m
}

func main() {
	text := []string{
		"Hello World !",
		"Ahahaha !",
	}

	mp := map[rune]int{}

	mpChannel := make(chan map[rune]int, len(text))

	for _, str := range text {
		go func(s string) {
			mpChannel <- GetFreq(s)
			fmt.Println("done:" + s)
		}(str)
	}

	for range text {
		m := <-mpChannel
		fmt.Println("got:", m)
		for k, v := range m {
			mp[k] += v
		}
	}

	fmt.Println(mp)
}
