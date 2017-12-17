package main

import (
	"fmt"
	"time"
)

var irc = make(chan string)
var sms = make(chan string)

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

//para o selects
func eai(canal chan string) {
	for {
		canal <- "\neai"
	}
}

func blz(canal chan string) {
	for {
		canal <- "\nblz"
	}
}

func impressora() {
	var msg string
	for {
		select {
		case msg = <-irc:
			fmt.Println(msg)
			time.Sleep(time.Second * 1)
		case msg = <-sms:
			fmt.Println("ho ho ho ho ", msg)
		}
		time.Sleep(time.Second * 1)
	}
}

func main() {
	//goroutines
	go pinger(irc)
	go eai(sms)
	go ponger(irc)
	go blz(sms)
	go impressora()

	var entrada string
	fmt.Scanln(&entrada)

}
