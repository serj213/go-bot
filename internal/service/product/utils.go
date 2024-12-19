package product

import (
	"fmt"
	"strconv"
	"strings"
)

func hasElemProduct(products []Product, target string) (Product, error){
	var result Product

	for idx, product := range products {
		if strconv.Itoa(idx) == strings.TrimSpace(target){
			result = product
			return result, nil
		}
	}
	return result, fmt.Errorf("doesn exist %s", target)
}