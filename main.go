package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func index(writer http.ResponseWriter, request *http.Request) {
	indexhtml := `
	hello travis-ci.
`
	fmt.Fprintln(writer, indexhtml)
}

func main() {
	// 整几个命令行参数
	port := flag.String("p", "8080", "监听端口")
	flag.Parse()

	fmt.Println("start listening at `http://localhost:" + *port + "/` ...")
	http.HandleFunc("/", index)

	err := http.ListenAndServe(":"+*port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
