package quote

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
)

type Quote struct {
	quote  string
	author string
}

func (q Quote) String() string {
	return fmt.Sprintf("%s\n    - %s", q.quote, q.author)
}

func GetQuote() (*Quote, error) {
	quotesFile, err := os.Open(os.Getenv("HOME") + "/.classup/quotes.csv")

	if err != nil {
		return nil, err
	}

	defer quotesFile.Close()

	r := csv.NewReader(quotesFile)

	// get random number between 1 and 1664
	rand := rand.Intn(1663) + 1
	for i := 0; i < rand; i++ {
		_, err = r.Read()

		if err != nil {
			return nil, err
		}
	}

	quoteCSV, err := r.Read()

	if err != nil {
		return nil, err
	}

	quote := Quote{quoteCSV[1], quoteCSV[0]}
	return &quote, nil
}
