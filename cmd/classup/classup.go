package main

import (
	"os"

	"github.com/syeadz/classup/internal/cli"
	coursedb "github.com/syeadz/classup/internal/database"
)

func main() {
	args := os.Args[1:]

	if len(args) > 0 {
		cli.Picker(args)
	} else {
		coursedb.PrintAllCourses()
	}

}
