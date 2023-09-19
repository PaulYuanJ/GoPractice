package server

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func EP_websocket() {

	http.HandleFunc("/ws/log", func(w http.ResponseWriter, r *http.Request) {
		conn, _ := upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity
		//
		//conn.SetPingHandler(func(appData string) error {
		//	fmt.Println("Receive Ping Msg..")
		//	err := conn.WriteControl(
		//		websocket.PongMessage,
		//		nil,
		//		time.Now().Add(2 * time.Second),
		//	)
		//	return err
		//})

		conn.SetPongHandler(func(appData string) error {
			fmt.Println("Receive Pong Msg..")
			return nil
		})

		conn.SetPingHandler(func(appData string) error {
			fmt.Println("Receive Ping Msg..")
			err := conn.WriteMessage(websocket.PingMessage, nil)
			return err
		})

		for {
			// Read message from browser
			msgType, msg, err := conn.ReadMessage()

			fmt.Println(msgType, msg, err)
			if err != nil {
				return
			}

			// Print the message to the console
			fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))

			// Write message back to browser
			if err = conn.WriteMessage(msgType, msg); err != nil {
				return
			}
		}
	})

	http.HandleFunc("/ws/metric", func(w http.ResponseWriter, r *http.Request) {
		conn, _ := upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity
		//
		//conn.SetPingHandler(func(appData string) error {
		//	fmt.Println("Receive Ping Msg..")
		//	err := conn.WriteControl(
		//		websocket.PongMessage,
		//		nil,
		//		time.Now().Add(2 * time.Second),
		//	)
		//	return err
		//})

		conn.SetPongHandler(func(appData string) error {
			fmt.Println("Receive Pong Msg..")
			return nil
		})

		conn.SetPingHandler(func(appData string) error {
			fmt.Println("Receive Ping Msg..")
			return nil
		})

		for {
			// Read message from browser
			msgType, msg, err := conn.ReadMessage()

			fmt.Println(msgType, msg, err)
			if err != nil {
				return
			}

			// Print the message to the console
			fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))

			// Write message back to browser
			if err = conn.WriteMessage(msgType, msg); err != nil {
				return
			}
		}
	})

	http.HandleFunc("/ws/task", func(w http.ResponseWriter, r *http.Request) {
		conn, _ := upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity
		//
		//conn.SetPingHandler(func(appData string) error {
		//	fmt.Println("Receive Ping Msg..")
		//	err := conn.WriteControl(
		//		websocket.PongMessage,
		//		nil,
		//		time.Now().Add(2 * time.Second),
		//	)
		//	return err
		//})

		conn.SetPongHandler(func(appData string) error {
			fmt.Println("Receive Pong Msg..")
			return nil
		})

		conn.SetPingHandler(func(appData string) error {
			fmt.Println("Receive Ping Msg..")
			return nil
		})

		for {
			// Read message from browser
			msgType, msg, err := conn.ReadMessage()

			fmt.Println(msgType, msg, err)
			if err != nil {
				return
			}

			// Print the message to the console
			fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))

			// Write message back to browser
			if err = conn.WriteMessage(msgType, msg); err != nil {
				return
			}
		}
	})

	http.ListenAndServe(":8080", nil)

}
