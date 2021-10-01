# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /blackjack

# Download required Go modules 
COPY go.mod ./
RUN go mod download

COPY *.go ./

RUN go build -o /blackjack-start

CMD [ "/blackjack-start" ]
