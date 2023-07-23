package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	// 创建一个新的路由器
	r := mux.NewRouter()

	// 定义处理函数
	r.HandleFunc("/", HomeHandler)

	// 启动服务器
	fmt.Println("Server listening on port 8080")
	http.ListenAndServe(":8080", r)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Go with gorilla/mux!")
}
