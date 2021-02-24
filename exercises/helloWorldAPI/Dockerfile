FROM golang:latest AS builder

WORKDIR /go/src/
COPY ./ ./
#RUN go build
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" ./...

FROM alpine:latest

RUN apk add --no-cache libc6-compat
COPY --from=builder --chown=1001:0 /go/src/api ./
EXPOSE 8080

CMD ["./api"]

