package main


import (
	"fmt"
	"net/http"
	"ghost/helper/internal/discovery"
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

	http.ServeFile(w, r, "./data/test.txt")
}

func main() {

	go discovery.RegisterService()

	go discovery.DiscoverPeers()

	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/asset", assetHandler)
	

	fmt.Println("Ghost helper running on :8080")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		panic(err)
	}
}