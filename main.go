package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/convee/blog/handler"
	"github.com/convee/goboot"
)

var (
	addr = *flag.String("addr", "127.0.0.1:8001", "http server")
)

func main() {
	flag.Parse()
	goboot.Run("./conf/config.toml")
	startHTTPServer(addr)

}

func startHTTPServer(addr string) {
	fmt.Println("http server starting....")
	http.HandleFunc("/home", home)

	// 设置静态目录
	fsh := http.FileServer(http.Dir("./static"))
	http.Handle("/", http.StripPrefix("/", fsh))

	http.ListenAndServe(addr, nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	handler.GetArticleList(w)
}
