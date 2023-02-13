package main

import (
	"os"

	"github.com/syeadz/classup/internal/cli"
	"github.com/syeadz/classup/internal/database/coursedb"
)

func main() {
	args := os.Args[1:]

	if len(args) > 0 {
		cli.Picker(args)
	} else {
		coursedb.PrintAllCourses()
	}

}
