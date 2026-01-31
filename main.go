/*
Copyright © 2026 Bagrov Konstantin
*/
package main

import (
	"fmt"
	"os"

	"github.com/KostyaBagr/duple-duple/cmd"
	"github.com/KostyaBagr/duple-duple/internal/config"
)

func main() {
	if err := config.ReadCfgFile(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("CLI has just started")
	cmd.Execute()
	fmt.Println("CLI has just shutted down")
}
