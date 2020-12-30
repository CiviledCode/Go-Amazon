package actions

import (
	"fmt"
	"github.com/ngs/go-amazon-product-advertising-api/amazon"
	"log"
	"strings"
)

type ProductLookup struct {
}

func (p *ProductLookup) ID() []string {
	return []string{"productlookup", "lookup", "top", "pl"}
}

func (p *ProductLookup) Execute(params []string, client *amazon.Client) error {
	res, err := client.ItemSearch(amazon.ItemSearchParameters{
		SearchIndex:    amazon.SearchIndexMusic,
		ResponseGroups: []amazon.ItemSearchResponseGroup{amazon.ItemSearchResponseGroupLarge},
		Keywords:       strings.Join(params, " "),
	}).Do()

	if err != nil {
		return err
	}

	log.Printf("Total results found %d\n\n", res.Items.TotalResults)

	for placement, item := range res.Items.Item {
		fmt.Printf(`#%d_________________________
[Title] %v
[Price] %v
[URL]   %v
`, placement, item.ItemAttributes.Title, item.ItemAttributes.ListPrice, item.DetailPageURL)
	}

	return nil
}
