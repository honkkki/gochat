package main

import (
	"fmt"
	"github.com/honkkki/gochat/server"
	"log"
	"net/http"
)

var (
	addr   = ":9999"
	banner = `
    ____              _____
   |    |    |   /\     |
   |    |____|  /  \    | 
   |    |    | /----\   |
   |____|    |/      \  |

ChatRoom，start on：%s
`
)



func main()  {
	fmt.Printf(banner+"\n", addr)

	server.RegisterHandler()

	log.Fatal(http.ListenAndServe(addr, nil))
}