package models

import (
	"github.com/Lordjima/Covid-19-Fighter/back/config"
	"log"
	"time"
)

type Patient struct {
	Id          int       `json:"id"`
	PatientID   string    `json:"patientID"`
	PatientAddr string    `json:"patientAddr"`
	DecryptKey  string    `json:"decryptKey"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Patients []Patient

func NewPatient(p *Patient) {
	if p == nil {
		log.Fatal(p)
	}
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()

	err := config.Db().QueryRow("INSERT INTO patients (patientID, patientAddr, decryptKey, created_at, updated_at) VALUES ($1,$2,$3,$4,$5) RETURNING id;", p.PatientID, p.PatientAddr, p.DecryptKey, p.CreatedAt, p.UpdatedAt).Scan(&p.Id)

	if err != nil {
		log.Fatal(err)
	}
}

func FindPatientByPatientID(patientID string) *Patient {
	var patient Patient

	row := config.Db().QueryRow("SELECT * FROM patients WHERE patientID = $1;", patientID)
	err := row.Scan(&patient.Id, &patient.PatientID, &patient.PatientAddr, &patient.DecryptKey, &patient.CreatedAt, &patient.UpdatedAt)

	if err != nil {
		log.Fatal(err)
	}

	return &patient
}

func AllPatients() *Patients {
	var patients Patients

	rows, err := config.Db().Query("SELECT * FROM patients")

	if err != nil {
		log.Fatal(err)
	}

	// Close rows after all readed
	defer rows.Close()

	for rows.Next() {
		var p Patient

		err := rows.Scan(&p.Id, &p.PatientID, &p.PatientAddr, &p.DecryptKey, &p.CreatedAt, &p.UpdatedAt)

		if err != nil {
			log.Fatal(err)
		}

		patients = append(patients, p)
	}

	return &patients
}

func UpdatePatient(patient *Patient) {
	patient.UpdatedAt = time.Now()

	stmt, err := config.Db().Prepare("UPDATE patients SET patientID=$1, patientAddr=$2, decryptKey=$3, updated_at=$4 WHERE id=$5;")

	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(patient.PatientID, patient.PatientAddr, patient.DecryptKey, patient.UpdatedAt, patient.Id)

	if err != nil {
		log.Fatal(err)
	}
}

func DeletePatientByPatientID(patientID string) error {
	stmt, err := config.Db().Prepare("DELETE FROM patients WHERE patientID=$1;")

	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(patientID)

	return err
}
