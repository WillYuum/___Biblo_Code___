package models

import (
	"log"
	"strings"

	"github.com/gorilla/websocket"
)

func Reader(conn *websocket.Conn) {
	for {
		// read in a message
		_, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		if string(p) == "get_data" {
			InitializeData(conn)

		}
		if strings.Split(string(p), "")[0] == "h" {
			str := strings.Split(string(p), "historic,")
			GetHistoric(strings.Split(str[1], " - ")[0], strings.Split(str[1], " - ")[1], conn)

		}
		if strings.Split(string(p), "")[0] == "D" {
			str := strings.Split(string(p), "Dynamic,")
			GetDynamic(strings.Split(str[1], " - ")[0], strings.Split(str[1], " - ")[1], conn)

		}
		// print out that message for clarity

	}
}
