FROM golang:1.15-alpine as builder

WORKDIR /build

ARG VERSION
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY="https://proxy.golang.org,direct"

RUN apk add --no-cache git


COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .

RUN go build -a -ldflags "-w -X 'main.Version=$VERSION'" -tags 'netgo osusergo' -o /go/bin/foaas main.go
RUN ldd /go/bin/foaas 2>&1 | grep -q 'Not a valid dynamic program'

LABEL description=foaas
LABEL builder=true


FROM alpine
COPY --from=builder go/bin/foaas /usr/local/bin

WORKDIR usr/local/bin
ENTRYPOINT [ "foaas", "serve"]
EXPOSE 8080
