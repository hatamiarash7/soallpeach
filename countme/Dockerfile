FROM golang:1.22.1 AS builder

WORKDIR /

COPY go.mod go.sum ./

RUN go mod download

COPY main.go ./

RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o countme main.go

FROM alpine:3.16.2

COPY --from=builder /countme .

ENV PORT 80

ENV GOGC off

EXPOSE 80

ENTRYPOINT ["/countme"]
