package main

import (
	"fmt"
	"time"
)

func pinger(canal chan string) {
	//direção da seta é a direção da comunicação
	for {
		canal <- "ping"
	}
}

func ponger(canal chan string) {
	for {
		canal <- "pong"
	}
}

func impressora(canal chan string) {
	for {
		msg := <-canal
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	var canal chan string
	canal = make(chan string)
	//goroutines
	go pinger(canal)
	go ponger(canal)
	go impressora(canal)
	var entrada string
	fmt.Scanln(&entrada)

}
