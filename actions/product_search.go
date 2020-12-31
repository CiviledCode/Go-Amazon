package actions

import (
	"encoding/json"
	"fmt"
	"github.com/ngs/go-amazon-product-advertising-api/amazon"
	"io"
	"log"
	"strconv"
	"strings"
)

// ProductSearch allows us to search for products using keywords/names
// You cannot search for more than 32,767 products
type ProductSearch struct {
}

func (p *ProductSearch) ID() []string {
	return []string{"productsearch", "search", "top", "pl"}
}

func (p *ProductSearch) Usage() string {
	return "search [name] (resultsamount)"
}

func (p *ProductSearch) Execute(params []string, client *amazon.Client, writer io.Writer) error {
	log.SetOutput(writer)

	if len(params) < 1 {
		return fmt.Errorf("expected params length of %s, got %d. Usage: %s", "> 1", len(params), p.Usage())
	}

	res, err := client.ItemSearch(amazon.ItemSearchParameters{
		SearchIndex:    amazon.SearchIndexMusic,
		ResponseGroups: []amazon.ItemSearchResponseGroup{amazon.ItemSearchResponseGroupLarge},
		Keywords:       strings.Join(params, " "),
	}).Do()

	if err != nil {
		return err
	}

	var items []*ProductResult
	resultAmount, isAmount := strconv.Atoi(params[len(params)-1])
	var item amazon.Item
	var placement uint16
	if isAmount != nil {
		items = make([]*ProductResult, res.Items.TotalResults)
		for ph, item := range res.Items.Item {
			placement = uint16(ph)
			if placement >= 65534 {
				break
			}
			p, e := strconv.ParseFloat(item.ItemAttributes.ListPrice.Amount, 32)
			if e != nil {
				return e
			}
			items = append(items, &ProductResult{URL: item.DetailPageURL, Placement: placement, Price: float32(p)})
		}
	} else {
		items = make([]*ProductResult, resultAmount)
		for i := uint16(0); i < uint16(resultAmount); i++ {
			if i < uint16(len(res.Items.Item)) {
				item = res.Items.Item[i]
				placement = i
			}
			p, e := strconv.ParseFloat(item.ItemAttributes.ListPrice.Amount, 32)
			if e != nil {
				return e
			}
			items = append(items, &ProductResult{URL: item.DetailPageURL, Placement: placement, Price: float32(p)})
		}
	}

	for _, o := range items {
		j, _ := json.Marshal(o)
		log.Println(string(j))
	}

	return nil
}

type ProductResult struct {
	URL string `json:"url"`

	Placement uint16 `json:"place"`

	Price float32 `json:"price"`
}
