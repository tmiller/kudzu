package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"time"
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
	addrs := getAddrsEth0()
	fmt.Println(w, time.Now().Format(time.UnixDate))
	for _, addr := range addrs {
		fmt.Fprintln(w, addr.String())
	}
}

func main() {
	log.Println("Starting server on :80")
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":80", nil)

	if err != nil {
		log.Println(err)
	}
}
