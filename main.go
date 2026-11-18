package main

import (
	"log"
	"flag"
	// "net"
	// "net/http"
)

type MasterConfig struct {
	Endpoint // some network type
}

type ContainerConfig struct {
	Instances int
	// ...
}

type HTTPServer struct {
	MasterConfig
}



// implements a method to expose channels from which network requests are piped through
type NetworkInterface interface {
	SetupAndConnectMethod // should call master + store master details for other methods to access
	HeartbeatMethod() chan string // string chan as placeholder
	ContainerAlterationMethod() chan int
	WorkerAlterationMethod() chan string
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