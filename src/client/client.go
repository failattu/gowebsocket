package main

import (
	"log"
	"fmt"
    "golang.org/x/net/websocket"
	"os"
)
func main() {
	if len(os.Args) > 2 {
		origin := "http://localhost/"
		baseurl := os.Args[1]
		baseport := os.Args[2]
		url := "ws://"+ baseurl+":"+baseport+"/echo"
		ws, err := websocket.Dial(url, "", origin)
		if err != nil {
			log.Fatal(err)
		}
		if _, err := ws.Write([]byte("Red Hat  and lots of love for containers \n")); err != nil {
			log.Fatal(err)
		}
		var msg = make([]byte, 512)
		var n int
		if n, err = ws.Read(msg); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Received: %s.\n", msg[:n])
	}
}