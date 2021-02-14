# Testing Encrypt/Decrypt API (based on Gorilla Mux)

1. Import all dependencies with:
   `go get github.com/gorilla/mux`

   
2. To run the application and test API on the browser:
   `go build gorilla-mux-with-unit-tests`
   followed by
   `go run gorilla-mux-with-unit-tests`
   
   Then, test the encrypt API: localhost:8080/encrpyt/abcd

   Should return `400` HTTP code with message `Requested string length should be greater than 8 characters`
   
   Now try with: `localhost:8080/encrpyt/12345678`

Then, test the decrypt API: `localhost:8080/encrpyt/<encrypted string exactly as returned by the encrypt API>`

If you give wrong encrypted string, the API owuld return HTTP code `422` with an `Invalid Input` message.

3. To run all unit tests:
   `go test gorilla-mux-with-unit-tests`

   
  
