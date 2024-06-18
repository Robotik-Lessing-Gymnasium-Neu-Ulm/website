package config

import (
	"encoding/json"
	"log"
	"os"
	"strconv"
)

type CFG struct {
	Website struct {
		Host string `json:"Host"`
		Port int    `json:"Port"`
	} `json:"Website"`
}

func GetConfig() *CFG {
	if os.Args[1] == "prod" {
		var c CFG
		c.Website.Host = os.Getenv("Host")
		c.Website.Port, _ = strconv.Atoi(os.Getenv("Port"))
		return &c
	} else if os.Args[1] == "dev" {
		const file = "config/config.json"
		var config CFG

		cfgfile, err := os.Open(file)
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Fatalf("Error readeing config file: %d\n", err)
		}

		jsonParser := json.NewDecoder(cfgfile)
		jsonParser.Decode(&config)

		return &config
	} else {
		panic("Error: Wrong command line argument")
	}
}
