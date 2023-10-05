package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup

	//Quantidade de goroutines que serão executadas
	waitGroup.Add(2)

	go func() {
		escrever("Olá mundo!")
		waitGroup.Done() //tira um da contagem do waitGroup
	}()

	go func() {
		escrever("Programando em Go!")
		waitGroup.Done() //tira um da contagem do waitGroup
	}()

	waitGroup.Wait()

}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
