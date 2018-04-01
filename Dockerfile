# Builder image - BEGIN
FROM golang:alpine as builder

RUN apk --no-cache add curl git

WORKDIR /go/src/github.com/flameous/PotatoPartyBackend
# copy source files
COPY . .
RUN go get -d -v ./...
# building binary file
RUN go build -o app cmd/main.go
# Builder image - END


# Resulting image - BEGIN
FROM alpine:latest
WORKDIR /app/

# copying binary from builder image
COPY --from=builder /go/src/github.com/flameous/PotatoPartyBackend/app .
RUN chmod +x app
CMD [ "./app"]
# Resulting image - END