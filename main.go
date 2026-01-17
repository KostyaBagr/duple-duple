/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"log"
	"os"

	"github.com/KostyaBagr/duple-duple/cmd"
)

func main() {
	LOG_FILE := "/var/log/duple-duple.log"
	logFile, err := os.OpenFile(LOG_FILE, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Panic(err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)

	log.SetFlags(log.Lshortfile | log.LstdFlags)

	log.Println("CLI has just started")
	cmd.Execute()
	log.Println("CLI has just shutted down")
}
