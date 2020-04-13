FROM golang:1.13-alpine AS builder

RUN apk update && apk add --no-cache git

WORKDIR $GOPATH/src/mypackage/server/
COPY . .

RUN go get -d -v

RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/server


FROM alpine

COPY --from=builder /go/bin/server /go/bin/server

ENTRYPOINT ["/go/bin/server"]
