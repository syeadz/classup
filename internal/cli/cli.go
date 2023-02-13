package cli

import (
	"fmt"

	"github.com/syeadz/classup/internal/quote"
)

// TODO: add flag handler
func Picker(args []string) {
	switch args[0] {
	case "quote":
		fmt.Println(quote.GetQuote().String())
	}
}
