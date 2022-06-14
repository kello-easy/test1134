package main

import (
	"fmt"
	"net/http"
	"time"
)

func ping(w http.ResponseWriter, req *http.Request) {
	fmt.Println("pong")
}

func main() {
	fmt.Println("IM HERE!")
	time.Sleep(time.Second*5)
	http.HandleFunc("/api/user/register", ping)
	http.ListenAndServe(":8080", nil)
}
