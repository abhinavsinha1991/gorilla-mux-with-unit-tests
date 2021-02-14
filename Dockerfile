FROM golang:1.15

WORKDIR /
COPY . .
RUN go get -d github.com/gorilla/mux

CMD ["go","run","gorilla-mux-with-unit-tests"]