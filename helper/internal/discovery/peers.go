package discovery

import (
	"fmt"
	"ghost/helper/internal/transfer"
	"os"
	"time"

	"github.com/hashicorp/mdns"
)

func DiscoverPeers() {
	entriesCh := make(chan *mdns.ServiceEntry, 4)

	go func() {
		for entry := range entriesCh {
			host, _ := os.Hostname()

			if entry.Host == host+".local." {
				continue
			}

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

			go transfer.FetchAssets(peer.Addr, peer.Port)
			go transfer.FetchAsset(peer.Addr, peer.Port)
		}
	}()
	fmt.Println("Starting peer discovery...")
	for {
		fmt.Println("Scanning for peers...")
		mdns.Lookup("_ghost._tcp", entriesCh)

		time.Sleep(5 * time.Second)
	}
}
