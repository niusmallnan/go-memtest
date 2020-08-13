FROM ubuntu:18.04

RUN apt-get update && apt-get install -y git curl ca-certificates unzip xz-utils vim

COPY go-memtest /usr/bin/

CMD ["/usr/bin/go-memtest"]
