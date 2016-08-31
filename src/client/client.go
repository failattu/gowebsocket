package main

import (
	"log"
	"fmt"
    "golang.org/x/net/websocket"
	"os"
)
func socketcalls(remoteurl string, originurl string, message string){
	ws, err := websocket.Dial(remoteurl, "", originurl)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := ws.Write([]byte(message)); err != nil {
		log.Fatal(err)
	}
	var msg = make([]byte, 512)
	var n int
	if n, err = ws.Read(msg); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Received: %s.\n", msg[:n])
}
func main() {
	if len(os.Args) > 3 {
		origin := "http://localhost/"
		baseurl := os.Args[1]
		baseport := os.Args[2]
		message := os.Args[3]
		url := "ws://"+ baseurl+":"+baseport+"/echo"
		socketcalls(url,origin,message)
	}
}