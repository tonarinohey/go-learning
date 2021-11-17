package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type Message struct {
	Msg  string `json:"msg"`
	Name string `json:"name"`
}

var clients = make(map[*websocket.Conn]bool) // 接続されるクライアント
var broadcast = make(chan Message)           // メッセージ用ブロードキャストチャネル

// バッファサイズ指定する
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// websocketを接続するハンドラ
func websocketConnectHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil) // websocket通信を実現するためにハンドシェイクを行う
	if err != nil {
		log.Print(err)
		return
	}
	clients[conn] = true // 接続したクライアントを保存
}

// メッセージを送信するハンドラ
func messageHandler(w http.ResponseWriter, r *http.Request) {
	var msg Message
	msg.Msg = r.FormValue("msg")
	msg.Name = r.FormValue("name")
	broadcast <- msg // メッセージ用ブロードキャストチャネルに送信
}

func websocketMessages() {
	for {
		// チャネルからメッセージを受け取る
		msg := <-broadcast
		// 現在接続しているクライアントすべてにメッセージを送信する
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Println(err) // クライアントの接続が切れるとエラー
				client.Close()
				delete(clients, client)
			}
		}
	}
}

func main() {
	portNumber := "9000"
	http.Handle("/", http.FileServer(http.Dir("static")))
	http.HandleFunc("/ws", websocketConnectHandler)
	http.HandleFunc("/msg", messageHandler)
	log.Println("Server listening on port ", portNumber)
	go websocketMessages()
	http.ListenAndServe(":"+portNumber, nil)

}
