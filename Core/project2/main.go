package main

import (
	"fmt"
	"strings"
)

var productPrices = map[string]float64{
	"TSHIRT": 15.00,
	"MUG":    7.00,
	"HAT":    20.00,
	"BOOK":   26.00,
}

func calcItemPrice(item string) (float64, bool) {
	var finalPrice float64 = 0.00
	price, ok := productPrices[item]
	if !ok {
		// en el caso de que tenga el sufijo _SALE indicando oferta, se lo quita y se vuelve a evaluar
		if strings.HasSuffix(item, "_SALE") {
			originalCode := strings.TrimSuffix(item, "_SALE")
			price, ok = productPrices[originalCode]
			if ok {
				finalPrice = price * 0.80
				fmt.Printf("Item %s final price: %.2f \n", originalCode, finalPrice)
			}
		}
	} else {
		finalPrice = price
		fmt.Printf("Item %s final price: %.2f\n", item, finalPrice)
	}
	return finalPrice, ok
}

func main() {
	fmt.Println("--- Order Example ---")
	orderItems := []string{
		"TSHIRT",
		"TSHIRT",
		"HAT",
		"BOOK",
	}
	var totalPrice float64 = 0.0
	for _, item := range orderItems {
		price, ok := calcItemPrice(item)
		if ok {
			totalPrice += price
		}
	}
	fmt.Printf("Total price: %.2f\n", totalPrice)
}
