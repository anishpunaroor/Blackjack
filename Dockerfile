# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /blackjack

COPY go.sum ./

RUN go build -o /blackjack-start

CMD [ "/blackjack-start" ]
