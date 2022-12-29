package main

import (
	"fmt"
	"log"

	"github.com/fazilnbr/project-workey/pkg/config"
)

func main() {
	fmt.Println("starting my job-portal project in clean code architecture")

	fmt.Println("starting my job-portal project in clean code architecture")
	config, configErr := config.LoadConfig()
	if configErr != nil {
		log.Fatal("cannot load config: ", configErr)
	}
	fmt.Printf("\n\nconfig : %v\n\n", config)
}
