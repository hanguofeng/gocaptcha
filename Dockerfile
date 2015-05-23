FROM centos
RUN yum -y update
RUN yum install -y golang git mercurial memcached
RUN mkdir -p /home/work/gopath
ENV GOPATH /home/work/gopath;
RUN go get github.com/hanguofeng/gocaptcha/samples/gocaptcha-server;
WORKDIR $GOPATH/src/github.com/hanguofeng/gocaptcha/samples/gocaptcha-server
RUN ["go","build"]
EXPOSE 80
ENTRYPOINT ["./gocaptcha-server"]
