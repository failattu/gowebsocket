FROM golang
ADD ./ /go/src/myserver
RUN go get golang.org/x/net/websocket
RUN go install myserver
ENTRYPOINT /go/bin/myserver
EXPOSE 12345
