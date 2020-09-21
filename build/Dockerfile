FROM golang:alpine

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /go/src/agner.io/boast

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY ./build/boast.toml .
COPY . .

RUN apk update && apk upgrade && \
    apk add --no-cache make
RUN make test
RUN make

EXPOSE 53/udp 80 443 1337 8080 8443

CMD ["./boast"]
