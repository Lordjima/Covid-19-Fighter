package controllers

import (
	"encoding/json"
	"github.com/Lordjima/Covid-19-Fighter/back/models"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

func PatientCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
	}

	var patient models.Patient

	err = json.Unmarshal(body, &patient)

	if err != nil {
		log.Fatal(err)
	}

	models.NewPatient(&patient)

	json.NewEncoder(w).Encode(patient)
}

func PatientIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(models.AllPatients())
}

func PatientShow(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)

	patient := models.FindPatientByPatientID(vars["patientID"])

	json.NewEncoder(w).Encode(patient)
}

func PatientUpdate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
	}

	patient := models.FindPatientByPatientID(vars["patientID"])

	err = json.Unmarshal(body, &patient)

	models.UpdatePatient(patient)

	json.NewEncoder(w).Encode(patient)
}

func PatientDelete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)

	err := models.DeletePatientByPatientID(vars["id"])

	if err != nil {
		log.Fatal(err)
	}

}
