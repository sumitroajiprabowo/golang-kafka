FROM golang:alpine

RUN mkdir /app

WORKDIR /app

COPY . /app

ADD go.mod .
ADD go.sum .

RUN go mod download
