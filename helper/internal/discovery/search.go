package discovery

import (
	"encoding/json"
	"fmt"
	"net/http"

	"ghost/helper/internal/transfer"
)

func FindPeerWithAsset(
	name string,
) *Peer {

	for _, peer := range Peers {

		url := fmt.Sprintf(
			"http://%s:%d/assets",
			peer.Addr,
			peer.Port,
		)

		resp, err := http.Get(url)

		if err != nil {
			continue
		}

		var assets []transfer.Asset

		err = json.NewDecoder(resp.Body).Decode(&assets)

		resp.Body.Close()

		if err != nil {
			continue
		}

		for _, asset := range assets {

			if asset.Name == name {
				return &peer
			}
		}
	}

	return nil
}