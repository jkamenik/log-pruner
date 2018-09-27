# compile step
FROM golang:1.11-alpine3.8 as builder

ENV GOBIN=/usr/local/bin
ENV SRC_PATH=$HOME/go/src/github.com/jkamenik/log-pruner

RUN echo $SRC_PATH && mkdir -p $SRC_PATH && pwd && ls -laR

COPY . /$SRC_PATH/

RUN cd $SRC_PATH && ls -la && pwd && go install && ls -la $GOBIN


# Image step
FROM alpine:3.8

COPY --from=builder /usr/local/bin/log-pruner /usr/local/bin/log-pruner

ENTRYPOINT ["/usr/local/bin/log-pruner"]