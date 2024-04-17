package viscosity

import (
	"fmt"
	"log"
	"net/http"

	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	// URL of the website containing the table
	url := "https://www.michael-smith-engineers.co.uk/resources/useful-info/approximate-viscosities-of-common-liquids-by-type"

	// Send a GET request to the website
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Parse the HTML document
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the table with the desired data
	doc.Find("table tr").Each(func(i int, s *goquery.Selection) {
		// Extract the name of the liquid and its viscosity
		liquid := strings.TrimSpace(s.Find("td:nth-child(1)").Text())
		viscosity := strings.TrimSpace(s.Find("td:nth-child(2)").Text())

		// Display the liquid name and its viscosity
		fmt.Printf("%s: %s\n", liquid, viscosity)
	})

}

// Function to retrieve the viscosity of a specific liquid from the table
func getViscosity(doc *goquery.Document, liquidName string) string {
	var viscosity string

	doc.Find("table tr").Each(func(i int, s *goquery.Selection) {
		// Extract the name of the liquid and its viscosity
		liquid := strings.TrimSpace(s.Find("td:nth-child(1)").Text())
		if strings.EqualFold(liquid, liquidName) {
			viscosity = strings.TrimSpace(s.Find("td:nth-child(2)").Text())
		}
	})

	return viscosity
}
