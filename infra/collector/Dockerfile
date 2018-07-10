FROM golang:1.10.3
RUN go get -u github.com/golang/dep/cmd/dep && \
    go get -u github.com/jteeuwen/go-bindata/...

WORKDIR /go/src/github.com/Bo0km4n/d2d/infra/collector
COPY ["./Gopkg.toml", "./Gopkg.lock", "./"]
RUN dep ensure -v -vendor-only=true

COPY [".", "."]
WORKDIR /go/src/github.com/Bo0km4n/d2d/infra/collector
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o collector .
ENTRYPOINT ["./collector"]
CMD [""]
