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

	files, err := os.ReadDir("./data")

	if err != nil {
		http.Error(w, "Failed to read data directory", 500)
		return
	}

	var assets []transfer.Asset

	for _, file := range files {

		info, err := file.Info()

		if err != nil {
			continue
		}

		path := "./data/" + info.Name()

		hash, err := transfer.GenerateHash(path)

		if err != nil {
			continue
		}

		assets = append(assets, transfer.Asset{
			Name: info.Name(),
			Size: info.Size(),
			Hash: hash,
		})
	}

	json.NewEncoder(w).Encode(assets)
}

func requestHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

	enableCors(&w)

	name := r.URL.Query().Get("name")

	path := "./data/" + name

	_, err := os.Stat(path)

	if err == nil {

		fmt.Println("Serving local asset:", name)

		http.ServeFile(w, r, path)
		return
	}

	fmt.Println("Asset missing locally")

	peer := discovery.FindPeerWithAsset(name)

	if peer == nil {

		http.Error(
			w,
			"Asset not found",
			404,
		)

		return
	}

	fmt.Println(
		"Fetching from peer:",
		peer.Name,
	)

	err = transfer.FetchAsset(
		peer.Addr,
		peer.Port,
		name,
		"",
	)

	if err != nil {

		http.Error(
			w,
			"Peer fetch failed",
			500,
		)

		return
	}

	fmt.Println("Serving downloaded asset")

	http.ServeFile(
		w,
		r,
		path,
	)
}

func main() {

	go discovery.RegisterService()

	go discovery.DiscoverPeers()

	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/asset", assetHandler)
	http.HandleFunc("/assets", assetsHandler)
	http.HandleFunc("/request", requestHandler)

	fmt.Println("Ghost helper running on :8080")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		panic(err)
	}
}
