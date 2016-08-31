This is my Go Language and Websocket test.

You can build it on your own machine with 

##Quickguide##

go get golang.org/x/net/websocket

go build src/server

go build src/client

./server.exe

./client.exe

##Docker##

Will be composed later

cd src/server/

docker build -t gowebsocket .

docker run --publish 6060:12345 --name test --rm gowebsocket

##Windows 7##

vagrant up

vagrant ssh 

cd /vagrant/src/server/

docker build -t gowebsocket .

docker run --publish 6060:12345 --name test --rm gowebsocket
