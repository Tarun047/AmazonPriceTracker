package main

import (
	"AmazonPriceTracker/controller"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", controller.TrackPrice)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
