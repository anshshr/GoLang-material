package main

import (
	"fmt"
	"time"
)

// its helps us to achieve to paralleism means means multiple processes can run at he same time in te sysytem without conflicting each other


// just simply add the go keyword before the task it will execute as soon as the main is exexuted
// every function starting with the go will compete for the cpu and the function which will get the cpu get executed and gets out of the system

func SayHi(){
	fmt.Println("hiiiiiiii")
}

func sayHello(){
	fmt.Println("helloooooooo")
}

func main() {
	fmt.Println("learning about the go routines")

	go sayHello()
	go SayHi()
	go SayHi()
	go sayHello()
	go SayHi()
	go SayHi()
	go sayHello()
	go SayHi()
	go SayHi()

	time.Sleep(1000 * time.Millisecond)
}