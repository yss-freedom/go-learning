package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/url"
	"time"
)

func main() {
	interrupt := make(chan struct{})

	u := url.URL{Scheme: "ws", Host: "localhost:8081", Path: "/ws"}
	log.Printf("Connecting to %s", u.String())
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("Dial error: ", err)
		return
	}
	defer conn.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			select {
			case <-done:
				return
			case <-interrupt:
				err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
				if err != nil {
					log.Println("Write close error: ", err)
					return
				}
				select {
				case <-done:
				case <-time.After(time.Second):
				}
				return
			default:
				messageType, p, err := conn.ReadMessage()
				if err != nil {
					log.Println("Read error: ", err)
					return
				}
				log.Printf("Received: %s", p)
				if err := conn.WriteMessage(messageType, p); err != nil {
					log.Println("Write error: ", err)
					return
				}
			}
		}
	}()

	time.Sleep(10 * time.Second)
	close(interrupt)
	<-done
}
