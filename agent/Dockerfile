FROM golang:alpine AS builder

WORKDIR /build

ADD go.mod go.sum ./

COPY . .

RUN go build -o agent ./cmd/main.go

FROM alpine

WORKDIR /build

COPY --from=builder /build/agent /build/agent

CMD ["./agent"]