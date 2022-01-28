package main

import (
	"fmt"
	"ip-locator/app"
	"log"
	"os"
)

func main() {
	fmt.Println("Start point")

	app := app.Generate()
	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
