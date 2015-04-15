package main

import (
	"io"
	"net/http"
	"log"
	"os"
	"world"
)

// hello world, the web server
func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, world.Msg())
}

func main() {
	http.HandleFunc("/", HelloServer)
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
