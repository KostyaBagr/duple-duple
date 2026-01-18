/*
Copyright © 2026 Bagrov Konstantin
*/
package main

import (
	"log"
	"os"

	"github.com/KostyaBagr/duple-duple/cmd"
	"github.com/KostyaBagr/duple-duple/internal/config"
)

// Sets logging config
func setLoggingCfg() {
	LOG_FILE := "/var/log/duple-duple.log"
	logFile, err := os.OpenFile(LOG_FILE, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Panic(err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)

	log.SetFlags(log.Lshortfile | log.LstdFlags)

}

func main() {
	setLoggingCfg()

	cfg, err := config.ReadCfgFile()
	if err != nil {
		log.Fatal(err)
	}

	cmd.InitConfig(cfg)
	log.Println("CLI has just started")
	cmd.Execute()
	log.Println("CLI has just shutted down")
}
