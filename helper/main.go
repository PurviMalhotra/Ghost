package main

import (
	"encoding/json"
	"fmt"
	"ghost/helper/internal/discovery"
	"ghost/helper/internal/transfer"
	"net/http"
	"os"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	fmt.Fprintf(w, "Ghost helper alive")
}

func assetHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	name := r.URL.Query().Get("name")

	path := "./data/" + name

	http.ServeFile(w, r, path)
}

func assetsHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	w.Header().Set("Content-Type", "application/json")

	file, err := os.Stat("./data/bread.jpg")

	if err != nil {
		http.Error(w, "File not found", 404)
		return
	}

	assets := []transfer.Asset{
		{
			Name: file.Name(),
			Size: file.Size(),
		},
	}

	json.NewEncoder(w).Encode(assets)
}

func main() {

	go discovery.RegisterService()

	go discovery.DiscoverPeers()

	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/asset", assetHandler)
	http.HandleFunc("/assets", assetsHandler)

	fmt.Println("Ghost helper running on :8080")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		panic(err)
	}
}
