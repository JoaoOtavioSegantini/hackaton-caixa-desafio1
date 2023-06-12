FROM golang:1.20.4-alpine3.18 AS builder

RUN apk add --no-cache --update bash

WORKDIR /go/src

COPY . .

RUN go get -d -v ./framework/cmd/server \
  && go install -v ./framework/cmd/server \
  && go build -ldflags '-w -s' -a -installsuffix cgo -o server ./framework/cmd/server

FROM scratch

COPY --from=builder /go/src/server .
COPY .env .


CMD ["./server"]