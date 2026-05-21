package transfer

import (
	"fmt"
	"io"
	"net/http"
)

func FetchAssets(addr string, port int) {
	url := fmt.Sprintf(
		"http://%s:%d/assets",
		addr,
		port,
	)

	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("Failed to fetch assets:", err)
		return
	}

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	fmt.Println("Peer assets:", string(body))
}
