package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMyRouterAndHandler(t *testing.T) {
	req, _ := http.NewRequest("GET", "/encrypt/12345678", nil)
	res := httptest.NewRecorder()
	newEncryptDecryptServer().ServeHTTP(res, req)

	fmt.Println(res.Body.String())

	if res.Body.String() != "MTIzNDU2Nzg=" {
		t.Error("Expected MTIzNDU2Nzg= but got ", res.Body.String())
	}
}

func TestInvalidLengthRequest(t *testing.T) {
	req, _ := http.NewRequest("GET", "/encrypt/1234", nil)
	res := httptest.NewRecorder()
	newEncryptDecryptServer().ServeHTTP(res, req)

	fmt.Println(res.Body.String())

	if res.Body.String() != "Requested string length should be greater than 8 characters" && res.Code != 400 {
		t.Error("Expected MTIzNDU2Nzg= but got ", res.Body.String())
	}
}

func TestInvalidStringToBeDecrypted(t *testing.T) {
	req, _ := http.NewRequest("GET", "/decrypt/abcd111", nil)
	res := httptest.NewRecorder()
	newEncryptDecryptServer().ServeHTTP(res, req)

	fmt.Println(res.Body.String())

	if res.Body.String() != "Invalid Input" && res.Code != 422 {
		t.Error("Expected MTIzNDU2Nzg= but got ", res.Body.String())
	}
}

func TestNonExistentEndpoint(t *testing.T) {
	req, _ := http.NewRequest("GET", "/dummy", nil)
	res := httptest.NewRecorder()
	newEncryptDecryptServer().ServeHTTP(res, req)

	fmt.Println(res.Body.String())

	if res.Code != 404 {
		t.Error("Expected 404 but got ", res.Code)
	}
}
