package main

import (
	"fmt"
	"log"
	"os"

	"go.mod/app"
)

func main() {
	fmt.Printf("Initializing APP")
	aplicacao := app.Gerar()
	if err := aplicacao.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
