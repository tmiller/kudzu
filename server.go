package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"runtime"
	"sync"
	"time"
)

var (
	port          = flag.String("port", ":80", "Port to listen on")
	loadTime      = flag.Duration("load", 2*time.Minute, "Amount of time to apply load")
	reservedCores = flag.Int("reserve", 1, "Cores to avoid loading")
)

func getAddrsEth0() []net.Addr {
	interfaces, err := net.Interfaces()

	if err != nil {
		log.Fatal(err)
	}

	for _, item := range interfaces {
		if item.Name == "eth0" || item.Name == "en6" {
			addrs, err := item.Addrs()
			if err != nil {
				log.Fatal(err)
			}
			return addrs
		}
	}

	log.Fatal("No interface named 'eth0' found")
	return nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("Starting Request")
	for _, addr := range getAddrsEth0() {
		fmt.Fprintln(w, addr.String())
	}
}

var (
	underLoad = false
	loadMutex = &sync.Mutex{}
)

func loadHandler(w http.ResponseWriter, r *http.Request) {

	loadMutex.Lock()
	defer loadMutex.Unlock()

	if underLoad {
		fmt.Fprintln(w, "Already Under Load")
		return
	}

	fmt.Fprintln(w, "Starting Load")
	underLoad = true
	// load up CPU Cores
	if runtime.NumCPU()-*reservedCores == 0 {
		go doCPUWork()
		return
	}

	for i := 0; i < runtime.NumCPU()-*reservedCores; i++ {
		go doCPUWork()
	}
}

func main() {
	flag.Parse()
	log.Println("Starting server on", *port)
	http.HandleFunc("/", handler)
	http.HandleFunc("/load", loadHandler)
	err := http.ListenAndServe(*port, nil)

	if err != nil {
		log.Println(err)
	}
}
