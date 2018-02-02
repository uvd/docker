FROM golang:1.9

ENV CGO_ENABLED=0
ENV GOOS=linux

RUN go get -u -v github.com/heptio/authenticator/cmd/heptio-authenticator-aws
RUN go build -a -installsuffix cgo github.com/heptio/authenticator/cmd/heptio-authenticator-aws

FROM scratch
COPY --from=0 /go/heptio-authenticator-aws /heptio-authenticator-aws
ENTRYPOINT ["/heptio-authenticator-aws"]
