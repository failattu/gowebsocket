FROM golang
ADD ./ /go/src/myclient
RUN go get golang.org/x/net/websocket
RUN go install myclient
ENTRYPOINT /go/bin/myclient