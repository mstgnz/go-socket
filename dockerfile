FROM golang:1.20 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o gosocket ./cmd

FROM scratch as deploy
COPY --from=builder /app/gosocket /bin/app
ENTRYPOINT ["/bin/app"]