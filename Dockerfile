FROM golang:1.22-alpine as builder

LABEL authors="brunooliveira"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download
RUN go mod verify

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -v -o /app/main ./cmd

FROM alpine:latest
WORKDIR /root/

COPY --from=builder /app/main .

CMD ["./main"]