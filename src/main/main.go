package main

import (
	"fmt"
	ftp "github.com/jlaffaye/ftp"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/agent", agent)
	http.HandleFunc("/admin", admin)
	log.Fatalln(http.ListenAndServe("localhost:8000", nil))
}

func agent(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Method:\n%s\n", request.Method)
	fmt.Fprintf(writer, "URL:\n%s\n", request.URL)
	fmt.Fprintf(writer, "Header:\n%s\n", request.Header)
	fmt.Fprintf(writer, "ContentLength:\n%s\n", request.ContentLength)
	fmt.Fprintf(writer, "Form:\n%s\n", request.Form)
	fmt.Fprintf(writer, "MultipartForm:\n%s\n", request.MultipartForm)
	fmt.Fprintf(writer, "PostForm:\n%s\n", request.PostForm)
	fmt.Fprintf(writer, "Proto:\n%s\n", request.Proto)
	fmt.Fprintf(writer, "ProtoMajor:\n%s\n", request.ProtoMajor)
	fmt.Fprintf(writer, "ProtoMinor:\n%s\n", request.ProtoMinor)
	fmt.Fprintf(writer, "RemoteAddr:\n%s\n", request.RemoteAddr)
	fmt.Fprintf(writer, "RequestURI:\n%s\n", request.RequestURI)
	ftp, err := ftp.Connect("ftp://192.168.1.1:21")
	if nil != err {
		log.Fatalf("%v\n", err)
		fmt.Fprintf(writer, "%v\n", err)
	}
	ftp.RetrFrom(" ", 0)
}

func admin(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "%s %s %s\n", request.Host, request.Method, request.URL.Path)
}
