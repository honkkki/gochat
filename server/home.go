package server

import (
	"encoding/json"
	"fmt"
	"github.com/honkkki/gochat/global"
	"github.com/honkkki/gochat/logic"
	"html/template"
	"net/http"
)

// HomeHandleFunc 渲染首页.
func HomeHandleFunc(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles(global.RootDir + "/template/home.html")
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

func userListHandleFunc(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	userList := logic.Broadcaster.GetUserList()
	data, err := json.Marshal(userList)
	if err != nil {
		fmt.Fprint(w, `[]`)
		return
	} else {
		fmt.Fprint(w, string(data))
		return
	}
}