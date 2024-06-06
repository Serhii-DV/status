package github

import (
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func TestGetProducts(t *testing.T) {
	html := `
<div class="components-section font-regular">
	<div class="components-container two-columns">
		<div class="components-container">
			<div class="component-container">
				<div class="component-inner-container status-green">
					<span class="name">Product Green 1</span>
				</div>
			</div>
			<div class="component-container">
				<div class="component-inner-container status-green">
					<span class="name">
						Product Green 2
					</span>
				</div>
			</div>
			<!-- Ignore product -->
			<div class="component-container" style="display: none;">
				<div class="component-inner-container status-green">
					<span class="name">Hidden product</span>
				</div>
			</div>
			<!-- Ignore product -->
			<div class="component-container">
				<div class="component-inner-container status-green">
					<span class="name">Visit www.githubstatus.com for more information</span>
				</div>
			</div>
		</div>
	</div>
</div>
	`

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		t.Fatalf("Error creating document: %s", err)
	}

	products := getProducts(doc)
	expected := []Product{
		{Name: "Product Green 1", Status: "Normal"},
		{Name: "Product Green 2", Status: "Normal"},
	}

	if len(*products) != len(expected) {
		t.Fatalf("Expected %d products, got %d", len(expected), len(*products))
	}

	for i, product := range *products {
		if product.Name != expected[i].Name || product.Status != expected[i].Status {
			t.Errorf("Expected product %d to be %+v, got %+v", i, expected[i], product)
		}
	}
}
