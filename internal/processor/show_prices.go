package processor

import (
	"buscador/internal/models"
	"fmt"
)

func ShowPriceAVG(priceChannel <-chan models.PriceDetail, done chan<- bool) {
	var totalPrice float64
	countPrices := 0.0
	for price := range priceChannel {
		totalPrice += price.Value
		countPrices++
		avgPrice := totalPrice / countPrices
		fmt.Printf("[%s] Preco da loja %s | R$ %.2f | Media Atual= %.2f \n", price.Timestamp.Format("02/01/2006 15:04:05"), price.StoreName, price.Value, avgPrice)
	}
	done <- true
}
