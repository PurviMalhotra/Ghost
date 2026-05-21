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

		go FetchAsset(addr, port, asset.Name)
	}
}

func FetchAsset(addr string, port int, name string) {
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
		return
	}

	defer resp.Body.Close()

	file, err := os.Create("./data/downloaded-" + name)

	if err != nil {
		fmt.Println("File creation failed:", err)
		return
	}

	defer file.Close()

	_, err = io.Copy(file, resp.Body)

	if err != nil {
		fmt.Println("File write failed:", err)
		return
	}

	fmt.Println("Asset downloaded successfully")
}
