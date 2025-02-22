FROM golang:alpine AS builder

WORKDIR /build

ADD go.mod go.sum ./

COPY . .

# RUN go mod download

# RUN go install github.com/swaggo/swag/cmd/swag@latest

# RUN swag init -g internal/application/application.go

RUN go build -o parser ./cmd/main.go

FROM alpine

WORKDIR /build

COPY --from=builder /build/parser /build/parser
COPY --from=builder /build/docs /build/docs

CMD ["./parser"]