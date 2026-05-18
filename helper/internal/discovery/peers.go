package discovery

import (
	"fmt"
	"time"

	"github.com/hashicorp/mdns"
)

func DiscoverPeers() {
	entriesCh := make(chan *mdns.ServiceEntry, 4)

	go func() {
	for entry := range entriesCh {

		peer := Peer{
			Name: entry.Name,
			Addr: entry.AddrV4.String(),
			Port: entry.Port,
		}

		Peers[entry.Name] = peer

		fmt.Printf(
			"Found peer: %s (%s:%d)\n",
			peer.Name,
			peer.Addr,
			peer.Port,
		)

		fmt.Println("Current peers:", Peers)
	}
}()

	for {
		mdns.Lookup("_ghost._tcp", entriesCh)

		time.Sleep(5 * time.Second)
	}
}