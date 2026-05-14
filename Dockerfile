FROM golang:1.26 AS builder

WORKDIR /app
COPY go.mod go.sum ./
COPY vendor ./vendor
COPY . .
RUN CGO_ENABLED=0 go build -mod=vendor -o /app/devyunmu.com ./cmd/server

FROM scratch

WORKDIR /

COPY --from=builder /app/devyunmu.com /devyunmu.com
COPY --from=builder /app/templates /templates
COPY --from=builder /app/static /static


EXPOSE 8080
ENTRYPOINT ["/devyunmu.com"]
