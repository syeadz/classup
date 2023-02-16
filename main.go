/*
Copyright © 2023 Syed Yead Zaman s.yead.zaman@gmail.com
*/
package main

import (
	"log"

	"github.com/syeadz/classup/cmd"
)

func main() {
	err := cmd.Execute()

	if err != nil {
		log.Fatal(err)
	}
}
