package transfer

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func FetchAssets(addr string, port int) {
	url := fmt.Sprintf(
		"http://%s:%d/assets",
		addr,
		port,
	)

	fmt.Println("Requesting:", url)

	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("Failed to fetch assets:", err)
		return
	}

	defer resp.Body.Close()

	var assets []Asset

	err = json.NewDecoder(resp.Body).Decode(&assets)

	if err != nil {
		fmt.Println("JSON decode failed:", err)
		return
	}

	fmt.Println("Peer assets:", assets)

	for _, asset := range assets {
		fmt.Println("Found asset:", asset.Name)

		if HasHash(asset.Hash) {

			fmt.Println(
				"Already cached:",
				asset.Name,
			)

			continue
		}

		go FetchAsset(
			addr,
			port,
			asset.Name,
			asset.Hash,
		)
	}
}

func FetchAsset(
	addr string,
	port int,
	name string,
	expectedHash string,
) error {
	url := fmt.Sprintf(
		"http://%s:%d/asset?name=%s",
		addr,
		port,
		name,
	)

	fmt.Println("Downloading asset from:", url)

	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("Download failed:", err)
		return err
	}

	defer resp.Body.Close()

	file, err := os.Create("./data/downloaded-" + name)

	if err != nil {
		fmt.Println("File creation failed:", err)
		return err
	}

	defer file.Close()

	_, err = io.Copy(file, resp.Body)

	if err != nil {
		fmt.Println("File write failed:", err)
		return err
	}

	hash, err := GenerateHash("./data/downloaded-" + name)

	if err != nil {
		fmt.Println("Hash verification failed")
		return err
	}

	fmt.Println("Downloaded file hash:", hash)

	if hash != expectedHash {
		fmt.Println("Hash mismatch detected")
		return err
	}

	fmt.Println("Hash verified successfully")

	fmt.Println("Asset downloaded successfully")

	return nil
}

func FindPeerWithAsset(name string) *Peer {

	for _, peer := range discovery.Peers {

		url := fmt.Sprintf(
			"http://%s:%d/assets",
			peer.Addr,
			peer.Port,
		)

		resp, err := http.Get(url)

		if err != nil {
			continue
		}

		var assets []Asset

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