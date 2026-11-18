package main

import (
	"log"
	"flag"
	"net/netip"
	// "net"
	// "net/http"
)

type ContainerConfig struct {
	Instances int
	// ...
}

type HTTPServer struct {
	MasterConfig
	ReqChan chan string
}

// change to one channel
func (s *HTTPServer) SetupAndConnectMethod() {
	// http call with s.masterAddr
	// create channel to pipe requests into
	// sets up all http handlers
	return 
}

func (s *HTTPServer) GetChan() (chan string) {
	return s.ReqChan
}

func New() (s *HTTPServer) {

}

// struct justs addr netip.Addr 
// implements a method to expose channels from which network requests are piped through
type NetworkInterface interface {
	SetupAndConnectMethod() // should call master 
	GetChan() chan string
} 



// connectToMaster
// startup HTTP Listener Services
func main() {
	masterAddr := *(flag.String("masterAddress", "localhost:6767", "ip of master machine"))
	masterPort := *(flag.Int("masterPort", 8080, "port of master machine"))

	flag.Parse()

	log.Printf("Instance started")
	log.Println(masterAddr, masterPort)

	// create httpserver
	// call httpserver.setup
}