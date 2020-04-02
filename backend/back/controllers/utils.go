package controllers

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type encryptBody struct {
	data       string
	passphrase string
}

func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

func Encrypt(data string, passphrase string) string {
	dataByte := []byte(data)
	block, _ := aes.NewCipher([]byte(createHash(passphrase)))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	ciphertext := gcm.Seal(nonce, nonce, dataByte, nil)
	return string(ciphertext)
}

func Decrypt(data string, passphrase string) string {
	dataByte := []byte(data)
	key := []byte(createHash(passphrase))
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := dataByte[:nonceSize], dataByte[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}
	return string(plaintext)
}

/*func PatientCreate(w http.ResponseWriter, r *http.Request) {
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
}*/

func EncryptEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
	}

	var encrypt encryptBody

	err = json.Unmarshal(body, &encrypt)

	if err != nil {
		log.Fatal(err)
	}

	encryptedData := Encrypt(encrypt.data, encrypt.passphrase)

	json.NewEncoder(w).Encode(encryptedData)
}

func DecryptEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
	}

	var encrypt encryptBody

	err = json.Unmarshal(body, &encrypt)

	if err != nil {
		log.Fatal(err)
	}

	decryptedData := Decrypt(encrypt.data, encrypt.passphrase)

	json.NewEncoder(w).Encode(decryptedData)
}
