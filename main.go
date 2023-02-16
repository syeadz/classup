/*
Copyright © 2023 Syed Yead Zaman s.yead.zaman@gmail.com
*/
package main

import (
	"os"

	"github.com/syeadz/classup/cmd"
)

func main() {
	err := cmd.Execute()

	if err != nil {
		os.Exit(-1)
	}
}
