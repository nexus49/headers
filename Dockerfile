FROM golang:1.10-alpine as builder

ARG DOCK_PKG_DIR=/go/src/github.com/nexus49/headers

RUN apk --no-cache --update add git

RUN mkdir -p $DOCK_PKG_DIR && mkdir -p bin

COPY ./ $DOCK_PKG_DIR
WORKDIR $DOCK_PKG_DIR

ENV GOOS=linux

RUN go get -u github.com/alecthomas/gometalinter && gometalinter --install && \
    gometalinter --vendor --deadline=2m --disable-all --enable=vet ./...
RUN go test --cover ./...
RUN go build -o ./main ./cmd

FROM alpine:3.8

COPY --from=builder /go/src/github.com/nexus49/headers/main .

CMD ["./main"]
