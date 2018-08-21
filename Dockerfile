FROM golang:1.10.3-alpine3.8 as builder

WORKDIR /go/src/app
COPY . .

RUN apk add --no-cache git
RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure
RUN go test
RUN go build -o boomvang

FROM golang:1.10.3-alpine3.8

WORKDIR /app/

COPY --from=builder /go/src/app/boomvang .

CMD ["/app/boomvang"]
