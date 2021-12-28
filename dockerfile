FROM golang:alpine AS builder

ENV GO111MODULE=on 
ENV CGO_ENABLED=0 
ENV GOOS=linux 
ENV GOARCH=amd64

WORKDIR /build

COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY ./ ./

RUN go build -o app .

WORKDIR /dist
RUN cp /build/app .


FROM alpine

WORKDIR /

COPY --from=builder /build/app /app
COPY config.json /
