FROM golang:1.24 AS builder
WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/dater/main.go

FROM scratch
COPY --from=builder ["/build/app", "/build/.env", "/"]
COPY --from=builder /build/static /static
EXPOSE 8000
CMD ["./app"]
