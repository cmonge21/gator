package main

import (
	"fmt"
	"log"

	"github.com/cmonge21/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}

	err = cfg.SetUser("Carly")
	if err != nil {
		log.Fatalf("Error setting user: %v", err)
	}

	updatedCfg, err := config.Read()
	if err != nil {
		log.Fatalf("Error reading updated config: %v", err)
	}

	fmt.Println(updatedCfg.DBURL)
}
