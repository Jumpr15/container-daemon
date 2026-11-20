package main

import (
	"container-daemon/types"

	"log"
)

type Container

type CDaemon struct {
	
	ReqChan chan any
}

func New(reqChan chan any) (*CDaemon) {
	daemon := CDaemon{
		ReqChan: reqChan,
	}
	return &daemon
}

func (d *CDaemon) HandleRequests() {
	for {
		val := <- d.ReqChan
		switch val.(type) { // now do cases
		case types.ContainerNum:
			log.Println("its a containernum(int)")
		case types.ContainerConfig:
			log.Println("its a containerconfig(struct)")
		case string:
			log.Println("its a string")
		
		case int:
			log.Println("its a integer")
		
		default:
			log.Println("unknown type")
		}
		log.Println(val)

	}
}

func main() {
	c := make(chan any)
	d := New(c)
	log.Println(c, d)
	go d.HandleRequests()
	c <- "Hello"
	c <- "World"
	c <- 12

	cn := types.ContainerNum(2)
	cc := types.ContainerConfig{}

	c <- cc
	c <- cn
}

// daemon should run seperate goroutine for looping forever and selecting via types