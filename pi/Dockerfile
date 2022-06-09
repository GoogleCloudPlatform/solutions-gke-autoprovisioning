FROM golang:1.18 as build-env
WORKDIR /go/src/app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY pi.go ./
RUN go build -o /go/bin/app

FROM gcr.io/distroless/base
COPY --from=build-env /go/bin/app /
ENTRYPOINT ["/app"]
