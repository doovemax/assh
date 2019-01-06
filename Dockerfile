FROM            golang:1.11-alpine as build
RUN             apk add --update --no-cache git gcc musl-dev make
COPY            go.* /go/src/github.com/doovemax/assh/
WORKDIR         /go/src/github.com/doovemax/assh
RUN             GO111MODULE=on go get .
COPY            . /go/src/github.com/doovemax/assh
RUN             make install

FROM            alpine
COPY            --from=build /go/bin/assh /bin/assh
ENTRYPOINT      ["/bin/assh"]
