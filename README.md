# Testing Encrypt/Decrypt API (based on Gorilla Mux)

1. Import all dependencies with:
   `go get github.com/gorilla/mux`

   
2. To run the application locally and test API on the browser:
   `go build gorilla-mux-with-unit-tests`
   followed by
   `go run gorilla-mux-with-unit-tests`

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Then, test the encrypt API: `localhost:8080/api/encrpyt/abcd`

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Should return `400` HTTP code with message `Requested string length should be greater than 8 characters`


&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Now try with: `localhost:8080/api/encrpyt/12345678`  

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Then, test the decrypt API: `localhost:8080/api/encrpyt/<encrypted string exactly as returned by the encrypt API>`

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;If you give wrong encrypted string, the API would return HTTP code `422` with an `Invalid Input` message.

3. To run all unit tests:
   `go test --cover`

4. To run the APi in docker:
   `docker build -t encrpy-decrypt-gorilla-api:v1 .`
   followed by
   `docker run -it -p 8080:8080 encrpy-decrypt-gorilla-api:v1`