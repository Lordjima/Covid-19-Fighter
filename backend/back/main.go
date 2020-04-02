package main

import (
	"github.com/Lordjima/Covid-19-Fighter/back/config"
	"github.com/Lordjima/Covid-19-Fighter/back/controllers"
	"github.com/Lordjima/Covid-19-Fighter/back/models"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	config.DatabaseInit()
	router := InitializeRouter()

	// Populate database
	models.NewPatient(&models.Patient{PatientID: "azerty1234", PatientAddr: "0x102aA24271a079f03c1c66295EFAB037c1DA151d", DecryptKey: "sport"})

	log.Fatal(http.ListenAndServe(":8080", router))
}

func InitializeRouter() *mux.Router {
	// StrictSlash is true => redirect /cars/ to /cars
	router := mux.NewRouter().StrictSlash(true)

	router.Methods("GET").Path("/patients").Name("Index").HandlerFunc(controllers.PatientIndex)
	router.Methods("POST").Path("/patient").Name("Create").HandlerFunc(controllers.PatientCreate)
	router.Methods("GET").Path("/patient/{id}").Name("Show").HandlerFunc(controllers.PatientShow)
	router.Methods("PUT").Path("/patient/{id}").Name("Update").HandlerFunc(controllers.PatientUpdate)
	router.Methods("DELETE").Path("/patient/{id}").Name("DELETE").HandlerFunc(controllers.PatientDelete)
	router.Methods("POST").Path("/encrypt").Name("Create").HandlerFunc(controllers.EncryptEndpoint)
	router.Methods("POST").Path("/decrypt").Name("Create").HandlerFunc(controllers.DecryptEndpoint)

	return router
}
