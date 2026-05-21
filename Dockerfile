FROM golang:1.26.2-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o /app/api cmd/api/main.go
RUN CGO_ENABLED=0 go build -o /app/migrator cmd/migrator/main.go

FROM alpine:3.22

WORKDIR /app

COPY --from=builder /app/api .
COPY --from=builder /app/migrator .

CMD ["./api"]