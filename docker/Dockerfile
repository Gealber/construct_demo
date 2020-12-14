# Antes de crear la imagen docker del posts hay q correr el comando 'go mod vendor'

FROM golang:1.13.3-alpine as builder

ENV CGO_ENABLED=0 GOFLAGS=-mod=vendor

WORKDIR /app/

# COPY <src> <dst>
COPY . ./

RUN go build -o /app/server ./main.go

LABEL maintainer="Gealber Morales Basalo <gealbermorales@gmail.com>"

# # new stage
FROM alpine:3.11.5

WORKDIR /app/
COPY --from=builder /app/server .

CMD ["./server" ]
