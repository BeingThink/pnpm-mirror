package server

import (
	"net/http"
	"github.com/gorilla/mux"
)


var version string
var arch string
var router *mux.Router

// https://github.com/pnpm/pnpm/releases/download/v8.7.6/pnpm-linux-arm64
func init() {
	router = mux.NewRouter()
	prefixRouter := router.PathPrefix("/pnpm/pnpm/releases/download/").Subrouter()
	prefixRouter.HandleFunc("/{version}/{arch}", func (w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		version = params["version"]
		arch = params["arch"]
		println(version, arch)
	})
	// TODO: 向客户端发送文件
}

func StartServer() {
	http.ListenAndServe(":8080", router)
}