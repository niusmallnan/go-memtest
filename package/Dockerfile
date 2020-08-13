FROM golang:1.14

RUN apt-get update && apt-get install -y git curl ca-certificates unzip xz-utils vim && \
    go get -u github.com/google/gops

COPY go-memtest /usr/bin/

CMD ["/usr/bin/go-memtest"]
