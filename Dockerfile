FROM golang:1.11.1
ENV GOPATH="/go"
RUN ["mkdir", "-p", "/go/src/github.com/rancher/demo-http"]
COPY * /go/src/github.com/rancher/demo-http/
WORKDIR /go/src/github.com/rancher/demo-http
RUN ["go", "build", "-o", "demo-http"]
CMD ["./demo-http"]