package server

import (
	"encoding/json"
	"github.com/gorilla/mux"
	pet "goAPI/pkg"
	"net/http"
)

type api struct{
	router http.Handler
	repository pet.PetRepository
}

type Server interface {
	Router() http.Handler
}

func New() Server  {
	a := &api{}

	r := mux.NewRouter()
	r.HandleFunc("/pets", a.fetchPets).Methods(http.MethodGet)
	r.HandleFunc("/pets/{ID:[a-zA-Z0-9_]+", a.fetchPet).Methods(http.MethodGet)
	a.router = r
	return a
}

func (a *api) Router() http.Handler {
	return a.router
}

func (a *api) fetchPets(writer http.ResponseWriter, request *http.Request) {
	pets, _ := a.repository.FetchPets()
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(pets)
}

func (a *api) fetchPet(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	pet, err := a.repository.FetchPetByID(vars["ID"])
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound) //Not found for sample simplicity
		json.NewEncoder(writer).Encode("Pet not found")
		return
	}
	json.NewEncoder(writer).Encode(pet)
}
