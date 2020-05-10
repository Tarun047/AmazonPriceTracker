package controller

import (
	"AmazonPriceTracker/model"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	GoQuery "github.com/PuerkitoBio/goquery"
)

//PriceTracker DTO for well price tracking

func makeRequest(url string) *http.Response {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Set("User-Agent", "Mozilla/5.0")
	fmt.Println("Looking up the current price at " + url)
	response, err := client.Do(request)
	if err != nil {
		panic(err.Error())
	}
	return response
}

func getTextFromSelector(document *GoQuery.Document, selector string) string {
	selection := document.Find(selector)
	return strings.TrimSpace(selection.Text())
}

func getPriceTrackerFromResponse(response *http.Response) model.PriceTracker {
	document, _ := GoQuery.NewDocumentFromResponse(response)
	priceTracker := model.PriceTracker{
		Title: getTextFromSelector(document, "#productTitle"),
		Price: getTextFromSelector(document, "#priceblock_ourprice"),
	}
	return priceTracker
}

//TrackPrice to track the price of a given product
func TrackPrice(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("product")
	response := makeRequest(url)
	defer response.Body.Close()
	priceTracker := getPriceTrackerFromResponse(response)
	json.NewEncoder(w).Encode(priceTracker)
}
