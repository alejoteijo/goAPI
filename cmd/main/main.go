package main

import (
	"flag"
	"fmt"
	"goAPI/cmd/sample-data"
	pet "goAPI/pkg"
	"goAPI/pkg/server"
	"goAPI/pkg/storage"
	"log"
	"net/http"
)

func main() {
	withData := flag.Bool("withData", false, "Initialize demo API with some pets")
	flag.Parse()
	var pets map[string]*pet.Pet

	if *withData {
		pets = sample.Pets
	}

	repo := storage.NewPetRepository(pets)
	s := server.New(repo)

	fmt.Println("The pet server is on tap now: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", s.Router()))
}
