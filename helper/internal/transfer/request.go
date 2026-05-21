package transfer

import (
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

	body, _ := io.ReadAll(resp.Body)

	fmt.Println("Peer assets:", string(body))
}

func FetchAsset(addr string, port int) {
	url := fmt.Sprintf(
		"http://%s:%d/asset",
		addr,
		port,
	)

	fmt.Println("Downloading asset from:", url)

	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("Download failed:", err)
		return
	}

	defer resp.Body.Close()

	file, err := os.Create("./data/downloaded-bread.jpg")

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
