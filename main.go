package main

import (
	"fmt"
	"log"
	"net/http"
	"wxcloudrun-golang/cmd"
	"wxcloudrun-golang/db"
	"wxcloudrun-golang/service"
)

func main() {
	if err := db.Init(); err != nil {
		panic(fmt.Sprintf("mysql init failed with %+v", err))
	}
	cmd.Init()
	http.HandleFunc("/", service.IndexHandler)
	log.Fatal(http.ListenAndServe(":80", nil))
}
