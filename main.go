package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	GoQuery "github.com/PuerkitoBio/goquery"
)

//PriceTracker DTO for well price tracking
type PriceTracker struct {
	Title string
	Price string
}

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

func getPriceTrackerFromResponse(response *http.Response) PriceTracker {
	document, _ := GoQuery.NewDocumentFromResponse(response)
	priceTracker := PriceTracker{
		Title: getTextFromSelector(document, "#productTitle"),
		Price: getTextFromSelector(document, "#priceblock_ourprice"),
	}
	return priceTracker
}

func main() {
	fmt.Println("Intializing ... ")
	if len(os.Args) != 2 {
		fmt.Println("Tool usage is AmazonPriceTracker <product_url>")
		return
	}
	response := makeRequest(os.Args[1])
	defer response.Body.Close()
	priceTracker := getPriceTrackerFromResponse(response)
	fmt.Printf("Title = %s\n", priceTracker.Title)
	fmt.Printf("Price = %s\n", priceTracker.Price)
	fmt.Println("Done")
}
