FROM golang:1.18.0-buster
WORKDIR /go/src

COPY ./go.* ./

RUN go mod download

RUN go install github.com/cosmtrek/air@v1.40.4 \
  && go install github.com/99designs/gqlgen@v0.17.13 

COPY . ./

