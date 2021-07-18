package server

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var RootDir string

func RegisterHandler() {
	getRootDir()

	// 广播消息处理
	//go logic.Broadcaster.Start()

	http.HandleFunc("/", HomeHandleFunc)
	http.HandleFunc("/ws", WebSocketHandleFunc)
}

func getRootDir() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal("get work dir fail: ", err)
	}

	var infer func(d string) string
	infer = func(d string) string {
		if exists(d + "/cmd") {
			return d
		}

		return infer(filepath.Dir(d))
	}

	RootDir = infer(wd)
}

func exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}


