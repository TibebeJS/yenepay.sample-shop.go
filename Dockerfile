#Compile stage
FROM golang:1.15.6-alpine3.12 AS build-env
ENV CGO_ENABLED 0
ENV PORT $PORT
ENV DATABASE_URL $DATABASE_URL
RUN apk add --no-cache git
RUN mkdir -p /go/src/github.com/TibebeJs

# Install revel framework
RUN go get github.com/revel/cmd/revel
RUN go get github.com/go-gorp/gorp

ADD . /go/src/github.com/TibebeJs/yenepay.sample-shop.go

#build revel app
WORKDIR /go/src/github.com/TibebeJs/yenepay.sample-shop.go
RUN rm -rf target
RUN revel build . dev

# Final stage
FROM golang:1.15.6-alpine3.12
ENV PORT $PORT
EXPOSE $PORT
WORKDIR /
COPY --from=build-env /go/src/github.com/TibebeJs/yenepay.sample-shop.go/dev /
ENTRYPOINT /run.sh