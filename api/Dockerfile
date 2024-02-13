FROM golang:1.20 as base

RUN apt-get update && apt-get -y install git

RUN mkdir /api

WORKDIR /api

COPY go.mod .

COPY go.sum .

RUN go mod download


FROM base as builder

COPY . .

RUN CGO_ENABLED=0 go build -o apiexecutable ./cmd/api/main.go


FROM scratch

COPY --from=builder /api/apiexecutable /apiexecutable

EXPOSE 8080

CMD [ "/apiexecutable" ]