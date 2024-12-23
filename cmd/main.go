package main

import (
	"buscador/internal/fetcher"
	"buscador/internal/models"
	"buscador/internal/processor"
	"fmt"
	"time"
)

func main() {

	start := time.Now()
	//buffer
	priceChennel := make(chan models.PriceDetail, 4)
	//len(qunatidade atual) e cap(tamanho max)
	done := make(chan bool)

	go processor.ShowPriceAVG(priceChennel, done)
	go fetcher.FechPrices(priceChennel)

	<-done //atua como um sincronizador

	fmt.Printf("\n Tempo total= %s", time.Since(start))
}
