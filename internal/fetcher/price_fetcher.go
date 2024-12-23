package fetcher

import (
	"buscador/internal/models"
	"math/rand"
	"sync"
	"time"
)

func FechPrices(priceChannel chan<- models.PriceDetail) {

	var wg sync.WaitGroup

	wg.Add(4)

	go func() {
		defer wg.Done()
		priceChannel <- FetchPriceFromSite1()
	}()

	go func() {
		defer wg.Done()
		priceChannel <- FetchPriceFromSite2()
	}()

	go func() {
		defer wg.Done()
		priceChannel <- FetchPriceFromSite3()
	}()

	go func() {
		defer wg.Done()
		FetchAndSendMultiplesPrices(priceChannel)
	}()

	wg.Wait()
	close(priceChannel)
}

// Buscar precos de diferentes sites

func FetchPriceFromSite1() models.PriceDetail {
	time.Sleep(1 * time.Second)
	return models.PriceDetail{StoreName: "Americanas", Value: rand.Float64() * 100, Timestamp: time.Now()}
}

func FetchPriceFromSite2() models.PriceDetail {
	time.Sleep(3 * time.Second)
	return models.PriceDetail{StoreName: "Magazine Luiza", Value: rand.Float64() * 100, Timestamp: time.Now()}
}

func FetchPriceFromSite3() models.PriceDetail {
	time.Sleep(2 * time.Second)
	return models.PriceDetail{StoreName: "Casas Bahia", Value: rand.Float64() * 100, Timestamp: time.Now()}
}

func FetchAndSendMultiplesPrices(priceChannel chan<- models.PriceDetail) {
	time.Sleep(6 * time.Second)
	prices := []float64{
		rand.Float64() * 100,
		rand.Float64() * 100,
		rand.Float64() * 100,
		rand.Float64() * 100,
	}
	for _, price := range prices {
		priceChannel <- models.PriceDetail{
			StoreName: "Kabum",
			Value:     price,
			Timestamp: time.Now(),
		}
	}
}
