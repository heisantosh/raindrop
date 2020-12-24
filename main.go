package main

import (
	"fmt"
	"os"

	"github.com/hashicorp/mdns"
)

func main() {
	host, _ := os.Hostname()
	info := []string{"Groovy Gorilla"}
	service, _ := mdns.NewMDNSService(host, "_raindrop._tcp", "", "", 8000, nil, info)
	server, _ := mdns.NewServer(&mdns.Config{Zone: service})
	defer server.Shutdown()

	entriesCh := make(chan *mdns.ServiceEntry, 4)
	go func() {
		for entry := range entriesCh {
			fmt.Printf("%+v\n", entry)
		}
	}()

	mdns.Lookup("_raindrop._tcp", entriesCh)
	close(entriesCh)
}
