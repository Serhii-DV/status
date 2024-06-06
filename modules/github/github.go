package github

import (
	"fmt"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Product struct {
	Name   string
	Status string
}

func Run() {
	url := "https://www.githubstatus.com/"
	doc, err := goquery.NewDocument(url)

	if err != nil {
		log.Fatal(err)
	}

	products := getProducts(doc)

	// Output the extracted data
	if len(*products) == 0 {
		fmt.Println("No components found. Please check the HTML structure and selectors.")
	} else {
		for _, component := range *products {
			fmt.Printf("Product: %s, Status: %s\n", component.Name, component.Status)
		}
	}
}

func getProducts(doc *goquery.Document) *[]Product {
	var products []Product

	classToStatus := map[string]string{
		"status-green":  "Normal",
		"status-yellow": "Degraded",
		"status-orange": "Degraded",
		"status-red":    "Incident",
		"status-blue":   "Maintenance",
	}

	doc.Find(".components-section .component-inner-container").Each(func(i int, s *goquery.Selection) {
		if isElementHidden(s.Parent()) {
			log.Println("Component container element is hidden")
			return
		}

		// debugElement(s)

		name := strings.TrimSpace(s.Find(".name").Text())

		if name == "Visit www.githubstatus.com for more information" {
			return
		}

		var status string
		for className, value := range classToStatus {
			if s.HasClass(className) {
				status = value
				break
			}
		}

		if status == "" {
			log.Printf("No matching status class found for %s\n", name)
			return
		}

		if name != "" && status != "" {
			products = append(products, Product{
				Name:   name,
				Status: status,
			})
		}
	})

	return &products
}

func isElementHidden(sel *goquery.Selection) bool {
	style, exists := sel.Attr("style")
	if !exists {
		return false // Element does not have a style attribute, so it's not hidden
	}

	return strings.Contains(style, "display: none")
}

func debugElement(el *goquery.Selection) {
	html, err := el.Parent().Html()
	if err != nil {
		log.Printf("Error getting HTML: %s\n", err)
		return
	}

	log.Printf("Element as HTML: %s\n", html)
}
