package main

import (
	"fmt"
	"ghost/helper/internal/discovery"
	"net/http"
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

	http.ServeFile(w, r, "./data/bread.jpg")
}

func assetsHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	fmt.Fprintf(w, "bread.jpg")
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
