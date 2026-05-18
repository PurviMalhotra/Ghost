package discovery

import (
	"fmt"
	"os"

	"github.com/hashicorp/mdns"
)

func RegisterService() {
	host, err := os.Hostname()

	if err != nil {
		panic(err)
	}

	service, err := mdns.NewMDNSService(
		host,
		"_ghost._tcp",
		"",
		"",
		8080,
		nil,
		[]string{"Ghost Node"},
	)

	if err != nil {
		panic(err)
	}

	server, err := mdns.NewServer(&mdns.Config{
		Zone: service,
	})

	if err != nil {
		panic(err)
	}

	fmt.Println("Ghost service registered")

	_ = server
}