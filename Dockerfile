FROM golang:alpine AS builder

WORKDIR /build

ADD go.mod go.sum ./

COPY . .

RUN go build -o parser ./cmd/main.go

FROM alpine

WORKDIR /build

COPY --from=builder /build/parser /build/parser

CMD ["./parser"]