package main

import (
	"fmt"
	"math/rand"
	"motivation/quotes"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().Unix())
	index := rand.Intn(len(quotes.Quotes))
	quote := quotes.Quotes[index]
	fmt.Printf("\n\033[1m%10s\033[0m\n\n", quote.Text)
	fmt.Printf("-- \033[4m%10s\033[0m\n\n", quote.Author)
}
