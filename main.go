package main

import (
	"fmt"
	"os"
	"time"

	"github.com/hashicorp/mdns"
)

const (
	_service = "_raindrop._tcp"
)

func main() {
	host, _ := os.Hostname()
	info := []string{"Groovy Gorilla"}
	service, _ := mdns.NewMDNSService(host, _service, "", "", 8000, nil, info)
	server, _ := mdns.NewServer(&mdns.Config{Zone: service})
	defer server.Shutdown()

	entriesCh := make(chan *mdns.ServiceEntry, 4)
	go func() {
		for entry := range entriesCh {
			fmt.Printf("%+v\n", entry)
		}
	}()

	mdns.Query(&mdns.QueryParam{
		Service:             _service,
		Domain:              "local",
		Timeout:             20 * time.Second,
		Entries:             entriesCh,
		WantUnicastResponse: true,
	})
	time.Sleep(20 * time.Second)
	close(entriesCh)
}
