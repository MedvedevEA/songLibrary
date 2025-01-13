package main

import (
	"log"
	"songLibrary/internal/appserver"
)

func main() {
	if err := appserver.Run(); err != nil {
		log.Fatal(err)
	}

}
