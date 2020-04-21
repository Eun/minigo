FROM golang:1.14 as build-env

WORKDIR /go/src/minigo
ADD . /go/src/minigo

RUN go get -d -v ./...

RUN go build -o /go/bin/minigo

FROM gcr.io/distroless/base
COPY --from=build-env /go/bin/minigo /
CMD ["/minigo"]