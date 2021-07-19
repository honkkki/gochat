package server

import (
	"github.com/honkkki/gochat/logic"
	"net/http"
)

func RegisterHandler() {
	// 广播消息处理
	go logic.Broadcaster.Start()

	http.HandleFunc("/", HomeHandleFunc)
	http.HandleFunc("/ws", WebSocketHandleFunc)
	http.HandleFunc("/user_list", userListHandleFunc)
}


