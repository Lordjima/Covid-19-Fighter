package main

import (
	"github.com/Lordjima/Covid-19-Fighter/back/config"
	"github.com/Lordjima/Covid-19-Fighter/back/models"
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
