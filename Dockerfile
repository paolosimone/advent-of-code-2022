FROM golang:1.19

WORKDIR /code

RUN go install mvdan.cc/gofumpt@latest && \
  curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.50.1

ENTRYPOINT ["make", "run"]