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
}

// change to one channel
func (s *HTTPServer) SetupAndConnectMethod() {
	// http call with s.masterAddr
	// creates channels for servers to pipe requests into
	// sets up all servers 
	// stores all corresponding channel details for later retrieval
}

func (s *HTTPServer) GetHeartbeatChan() (chan string) {

}

func (s *HTTPServer) GetContainerConfigChan() (chan int) {

}

func (s *HTTPServer) GetWorkerConfigChan() (chan string){

}
// struct justs addr netip.Addr 
// implements a method to expose channels from which network requests are piped through
type NetworkInterface interface {
	SetupAndConnectMethod() // should call master 
	GetHeartbeatChan() chan string // string chan as placeholder
	GetContainerConfigChan() chan int
	GetWorkerConfigChan() chan string
} 
// all take in required network values to create network listeners/servers => return channels where values are piped through
// disaggregate between communication type: http/rpc etc.


// connectToMaster
// startup HTTP Listener Services
func main() {
	masterAddr := *(flag.String("masterAddress", "localhost:6767", "ip of master machine"))
	masterPort := *(flag.Int("masterPort", 8080, "port of master machine"))

	flag.Parse()

	log.Printf("Instance started")
	log.Println(masterAddr, masterPort)
}