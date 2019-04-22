FROM golang:1.12 AS build

RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
WORKDIR /go/src/github.com/jossemargt/cms-sao
COPY . .
RUN dep ensure \
    && CGO_ENABLED=0 GOOS=linux go build -v

FROM alpine:3.9

COPY --from=build /go/src/github.com/jossemargt/cms-sao/cms-sao /opt/cms-sao
RUN chmod ugo+x /opt/cms-sao
USER nobody

CMD ["/opt/cms-sao"]