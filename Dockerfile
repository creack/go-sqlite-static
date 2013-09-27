FROM		ubuntu:12.10
MAINTAINER	Guillaume J. Charmes <guillaume@dotcloud.com>

RUN		apt-get update
RUN		apt-get install -y curl mercurial build-essential
RUN		curl -s https://go.googlecode.com/files/go1.1.2.linux-amd64.tar.gz | tar -v -C /usr/local -xz

ENV		PATH 	/usr/local/go/bin:$PATH
ENV		GOPATH	/go
ADD		test.go	/src/
RUN		cd /src && go get -d && go build -ldflags '-linkmode external -extldflags -static -w' -o test .
