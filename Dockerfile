FROM golang:1.15.3 as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

ENV GO111MODULE=on

RUN CGO_ENABLED=0 GOOS=linux go build -o app

FROM alpine:3.14.0 as release

WORKDIR /app

RUN apk add tzdata

ENV ZONEINFO=/usr/share/zoneinfo

COPY --from=builder /app/app /app/

CMD ["./app"]