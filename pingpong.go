//canal: modo duas goroutines se comunicarem e sincronizarem sua execucao

package main

import (
	"fmt"
	"time"
)

func pingar(c chan string) { //palavra reservada para canal: chan
	for i := 0; ; i++ {
		c <- "ping" //usado para enviar e receber mensagem pelo canal
	}
}
func pongar(c chan string) { //palavra reservada para canal: chan
	for i := 0; ; i++ {
		c <- "pong" //usado para enviar e receber mensagem pelo canal
	}
}

func imprimir(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	var c chan string = make(chan string)

	go pingar(c)
	go imprimir(c)
	go pongar(c)

	var entrada string
	fmt.Scanln(&entrada)

}
