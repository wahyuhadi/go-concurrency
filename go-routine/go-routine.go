package main

import (
	"fmt"
	"time"
)

func main() {
	go count("test")
	count1("lolll")
}


func count (num string) {
	for i := 1; true; i++ {
		fmt.Println(i, num)
		time.Sleep(time.Millisecond * 500)
	}
}

func count1 (num string) {
	for i := 1; true; i++ {
		fmt.Println(i, num)
		time.Sleep(time.Millisecond * 500)
	}
}