package server

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var rootDir string

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

	rootDir = infer(wd)
}

func exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

func HomeHandleFunc(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles(rootDir + "/template/home.html")
	if err != nil {
		fmt.Fprint(w, "模板解析错误！")
		return
	}

	err = tpl.Execute(w, nil)
	if err != nil {
		fmt.Fprint(w, "模板执行错误！")
		return
	}
}

func WebSocketHandleFunc(w http.ResponseWriter, r *http.Request) {

}
