package main

import (
	"log"

	exampleusage "github.com/ayoseun/geth-lte/example_usage"
	"github.com/joho/godotenv"
)



func main(){if err := godotenv.Load(); err != nil {
	log.Println("No .env file found")
}
	exampleusage.Init()
	exampleusage.ContractmemPool()
}