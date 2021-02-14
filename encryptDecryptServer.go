package main

import (
	b64 "encoding/base64"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func newEncryptDecryptServer() http.Handler {
	r := mux.NewRouter()
	r.Handle("/encrypt/{normalstring}", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(encryptHandler)))
	r.Handle("/decrypt/{encryptedstring}", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(decryptHandler)))
	return r
}

func encryptHandler(w http.ResponseWriter, r *http.Request) {
	name, exists := mux.Vars(r)["normalstring"]

	if len(name) < 8 {
		w.WriteHeader(400)
		w.Write([]byte("Requested string length should be greater than 8 characters"))
		return
	}

	if !exists {
		w.WriteHeader(404)
		w.Write([]byte("Invalid endpoint"))
		return
	}

	w.Write([]byte(b64.StdEncoding.EncodeToString([]byte(name))))
}

func decryptHandler(w http.ResponseWriter, r *http.Request) {
	encryptedstring, exists := mux.Vars(r)["encryptedstring"]

	if !exists {
		w.WriteHeader(404)
		w.Write([]byte("Invalid endpoint"))
		return
	}

	a, b := b64.StdEncoding.DecodeString(encryptedstring)

	if b != nil {
		w.WriteHeader(422)
		w.Write([]byte("Invalid input"))
		fmt.Println(b.Error() + "\n String was: " + encryptedstring)
		return
	}

	w.Write(a)
}
