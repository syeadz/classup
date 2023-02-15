package quote

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
)

type Quote struct {
	quote  string
	author string
}

func (q Quote) String() string {
	return fmt.Sprintf("%s\n\t- %s", q.quote, q.author)
}

func GetQuote() Quote {
	quotesFile, err := os.Open(os.Getenv("HOME") + "/.classup/quotes.csv")

	if err != nil {
		fmt.Println(err)
		log.Fatalln("Coudln't open quotes file. Maybe deleted?")
	}

	defer quotesFile.Close()

	r := csv.NewReader(quotesFile)

	// get random number between 1 and 1664
	rand := rand.Intn(1663) + 1
	for i := 0; i < rand; i++ {
		_, err = r.Read()

		if err != nil {
			log.Fatalln("Error reading quotes.csv")
		}
	}

	quoteCSV, err := r.Read()

	if err != nil {
		log.Fatalln("Error reading quotes.csv")
	}

	quote := Quote{quoteCSV[1], quoteCSV[0]}
	return quote
}

func PrintQuote() {
	fmt.Println(GetQuote().String())
}
