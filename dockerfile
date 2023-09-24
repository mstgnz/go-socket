FROM golang:1.20 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN GOOS=linux CGO_ENABLED=0 go build -o goSocket ./cmd

FROM scratch as deploy
COPY --from=builder /app/goSocket /app/goSocket
COPY --from=builder /app/public /public
ENTRYPOINT ["/app/goSocket"]