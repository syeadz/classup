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
	genre  string
}

func GetQuote() Quote {
	quotesFile, err := os.Open("/etc/classup/quotes.csv")

	if err != nil {
		log.Fatalln("Coudln't open quotes file. Maybe deleted?")
	}

	defer quotesFile.Close()

	r := csv.NewReader(quotesFile)
	r.Comma = ';'

	rand := rand.Intn(75967)
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

	quote := Quote{quoteCSV[0], quoteCSV[1], quoteCSV[2]}
	return quote
}

func (q Quote) String() string {
	return fmt.Sprintf("%s\n\t- %s (%s)", q.quote, q.author, q.genre)
}
