/*
Copyright © 2023 Syed Yead Zaman s.yead.zaman@gmail.com
*/
package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/syeadz/classup/internal/quote"
)

// quoteCmd represents the quote command
var quoteCmd = &cobra.Command{
	Use:   "quote",
	Short: "Displays a motivational quote",
	Long: `'quote' displays a random motivational quote from a quotes.csv 
file containing 1663 different quotes along with their authors.`,

	RunE: func(cmd *cobra.Command, args []string) error {
		quote, err := quote.GetQuote()

		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("quotes.csv doesn't exist. Is it downloaded?")
			return nil
		}
		if err != nil {
			return err
		}

		fmt.Println(quote.String())
		return nil
	},
}

func init() {
	rootCmd.AddCommand(quoteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// quoteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// quoteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
