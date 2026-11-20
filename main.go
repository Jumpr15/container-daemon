package main

import (
	"container-daemon/types"

	"fmt"
	"log"
	"flag"
	"encoding/json"
	"strconv"
	// "net/netip"
	// "net"
	"net/http"
)

type HTTPServer struct {
	MasterAddress string
	ListenerPort int
	ReqChan chan any
}

func (s *HTTPServer) SetupListeners() {
	// num containers
	http.HandleFunc("/containerNum", func(w http.ResponseWriter, r *http.Request) {
		var num types.ContainerNum

		err := json.NewDecoder(r.Body).Decode(&num)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		log.Println(num)
		
		// pass to channel
		s.ReqChan <- num
	})

	// set container config
	http.HandleFunc("/containerConfig", func(w http.ResponseWriter, r *http.Request) {
		var config types.ContainerConfig 

		err := json.NewDecoder(r.Body).Decode(&config)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		log.Println(config)

		// pass to channel
		s.ReqChan <- config
	})

	// health check
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		log.Println("heartbeat!!!")
	})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":" + strconv.Itoa(s.ListenerPort)), nil)) // chane port

}

func (s *HTTPServer) SetupAndConnectMethod() {
	// http call with s.masterAddr
	// create channel to pipe requests into
	// sets up all http handlers
	return 
}

func (s *HTTPServer) GetChan() (chan any) {
	return s.ReqChan
}

func New(masterAddress string) (*HTTPServer) {
	s := HTTPServer{
		MasterAddress: masterAddress,
		ListenerPort: 8080,
		ReqChan: make(chan any),
	}
	return &s
}

// implements a method to expose channels from which network requests are piped through
type NetworkInterface interface {
	SetupAndConnectMethod() // should call master 
	GetChan() chan string
} 

// startup HTTP Listener Services
func main() {
	masterAddr := *(flag.String("masterAddress", "localhost", "ip of master machine"))
	masterPort := *(flag.Int("masterPort", 8000, "port of master machine"))

	flag.Parse()

	log.Printf("Instance started")
	log.Println(masterAddr, masterPort)

	// currently only ipv4 handling ** very brittle
	masterAddress := fmt.Sprintf(masterAddr + ":" + strconv.Itoa(masterPort)) 
	log.Println(masterAddress)

	// create httpserver
	server := New(masterAddress)
	log.Println(server)
	log.Println(server.GetChan())
	// call httpserver.setup
	// create http listeners
	server.SetupListeners() // make goroutine
}