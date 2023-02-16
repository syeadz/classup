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
		// if resp.IsUserError() {
		// 	resp.Cmd.Println("")
		// 	resp.Cmd.Println(resp.Cmd.UsageString())
		// }
		os.Exit(-1)
	}
}
