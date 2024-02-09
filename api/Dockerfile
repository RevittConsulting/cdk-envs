FROM golang:alpine AS builder

RUN mkdir /build

ADD . /build/

WORKDIR /build

RUN go build -o main .

FROM alpine

COPY . /app/

COPY --from=builder /build/main /app/

WORKDIR /app

EXPOSE 8080

CMD ["./main"]