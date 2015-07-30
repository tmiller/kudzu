package main

import (
	"fmt"
	"net"
	"net/http"
	"time"
)

func getAddrsEth0() ([]net.Addr, error) {
	interfaces, err := net.Interfaces()

	if err != nil {
		return nil, err
	}

	for _, item := range interfaces {
		if item.Name == "eth0" || item.Name == "en6" {
			addrs, err := item.Addrs()
			return addrs, err
		}
	}
	return nil, fmt.Errorf("No interface named 'eth0' found")
}

func handler(w http.ResponseWriter, r *http.Request) {
	addrs, err := getAddrsEth0()
	if err != nil {
		fmt.Fprint(w, err)
	} else {
		fmt.Fprintln(w, time.Now().Format(time.UnixDate))
		for _, addr := range addrs {
			fmt.Fprintln(w, addr.String())
		}
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
