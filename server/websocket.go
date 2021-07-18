package server

import (
	"github.com/honkkki/gochat/logic"
	"log"
	"net/http"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

func WebSocketHandleFunc(w http.ResponseWriter, r *http.Request) {
	conn, err := websocket.Accept(w, r, &websocket.AcceptOptions{InsecureSkipVerify: true})
	if err != nil {
		log.Println("websocket accept error:", err)
		return
	}

	nickname := r.FormValue("nickname")
	if l := len(nickname); l < 1 || l > 30 {
		wsjson.Write(r.Context(), conn, logic.NewErrorMessage("昵称长度不正确，请检查"))
		conn.Close(websocket.StatusUnsupportedData, "illegal nickname")
		return
	}

	if !logic.Broadcaster.CanEnterRoom(nickname) {
		wsjson.Write(r.Context(), conn, logic.NewErrorMessage("该昵称已经已存在！"))
		conn.Close(websocket.StatusUnsupportedData, "nickname exists!")
		return
	}

	user := logic.NewUser(conn, "", nickname, r.RemoteAddr)
	// 开启向用户发送消息的goroutine. user get message.
	go user.SendMessage(r.Context())
	user.MessageChannel<-logic.NewWelcomeMessage(user)

	// 广播消息
	msg := logic.NewUserEnterMessage(user)
	logic.Broadcaster.Broadcast(msg)

	// add to map.
	logic.Broadcaster.UserEntering(user)
	log.Println("user:", user.NickName, "joins to chatroom")

	// get user send message
	err = user.ReceiveMessage(r.Context())

	// user left from chatroom
	logic.Broadcaster.UserLeaving(user)
	msg = logic.NewUserLeaveMessage(user)
	logic.Broadcaster.Broadcast(msg)
	log.Println("user:", user.NickName, "left from chatroom")

	if err != nil {
		log.Println("get message from client error:", err)
		conn.Close(websocket.StatusInternalError, "get message from client error")
	} else {
		conn.Close(websocket.StatusNormalClosure, "")
	}

}
