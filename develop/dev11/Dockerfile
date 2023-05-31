# Stage 1 - build executable in go container
FROM golang:latest as builder

WORKDIR $GOPATH/src/wbL2/develop/dev11/
COPY . .

RUN export CGO_ENABLED=0 && make build

# Stage 2 - build final image
FROM alpine:latest

# Copy our static executable
COPY --from=builder /go/src/wbL2/develop/dev11/bin/dev11 go/bin/dev11

# Run the binary.
ENTRYPOINT ["go/bin/dev11"]
