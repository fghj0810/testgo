package main

import (
	h1 "./api/auth"
	"./common"
	"encoding/json"
	"log"
	"net/http"
)

func errHandler(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	ret := make(map[string]interface{}, 5)
	ret["code"] = common.Code_InvalidAPI
	encoder.Encode(ret)
}

func main() {
	http.HandleFunc("/", errHandler)                        // 设置访问的路由
	http.HandleFunc("/api/auth/do_auth", h1.HandleFunction) // 设置访问的路由
	err := http.ListenAndServe(":9090", nil)                // 设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
