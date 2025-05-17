FROM golang:1.24.2 AS builder

WORKDIR /andreiolaru.site.bff
COPY . .
RUN go mod download
RUN go build -o main ./cmd

FROM debian:bookworm-slim
WORKDIR /andreiolaru.site.bff

COPY --from=builder /andreiolaru.site.bff/main .
COPY .env .env

EXPOSE 3001
CMD ["./main"]
