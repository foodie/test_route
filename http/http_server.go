package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	info("start")
	c1 := make(chan struct{})

	go runServer1()
	go runServer2()
	<-c1
}

func runServer1() {
	info("runServer1 start")
	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		info("handle s1 /hi")
		fmt.Fprintf(w, "%s", "hi s1")
	})
	http.ListenAndServe(":8081", nil)
}

type MyHandler struct {
}

func (MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	info("handle s2 /")
	fmt.Fprintf(w, "%s", "hi s2")
}

func runServer2() {
	info("runServer2 start")
	s := &http.Server{
		Addr:           ":8082",
		Handler:        &MyHandler{},
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

func info(str string) {
	log.Println(str)
}
