package main

import (
	"fmt"
	ftp "github.com/jlaffaye/ftp"
	"log"
	"net/http"
	"strings"
	"sync"
)

var mu sync.Mutex
var count int

const (
	method = "method"
	url    = "url"
	post   = "POST"
)

func main() {
	http.HandleFunc("/agent", agent)
	http.HandleFunc("/admin", admin)
	log.Fatalln(http.ListenAndServe("localhost:8000", nil))
}

func agent(writer http.ResponseWriter, request *http.Request) {
	//	fmt.Fprintf(writer, "Method:\n%s\n", request.Method)
	//	fmt.Fprintf(writer, "URL:\n%s\n", request.URL)
	//	fmt.Fprintf(writer, "Header:\n%s\n", request.Header)
	//	fmt.Fprintf(writer, "ContentLength:\n%s\n", request.ContentLength)
	//	fmt.Fprintf(writer, "Form:\n%s\n", request.Form)
	//	fmt.Fprintf(writer, "MultipartForm:\n%s\n", request.MultipartForm)
	//	fmt.Fprintf(writer, "PostForm:\n%s\n", request.PostForm)
	//	fmt.Fprintf(writer, "Proto:\n%s\n", request.Proto)
	//	fmt.Fprintf(writer, "ProtoMajor:\n%s\n", request.ProtoMajor)
	//	fmt.Fprintf(writer, "ProtoMinor:\n%s\n", request.ProtoMinor)
	//	fmt.Fprintf(writer, "RemoteAddr:\n%s\n", request.RemoteAddr)
	//	fmt.Fprintf(writer, "RequestURI:\n%s\n", request.RequestURI)

	if post != request.Method {
		log.Fatalf("accept wrong request %v\n only post was accepted", request)
		fmt.Fprintf(writer, "{\"code\":\"99\",\"desc\":\"wrong request %v\"\n", request)
		return
	}
	err := request.ParseForm()
	if nil != err {
		log.Fatalf("%v\n", err)
		fmt.Fprintf(writer, "{\"code\":\"99\",\"desc\":\"%v\"\n", err)
		return
	}
	m := request.Form.Get(method)
	u := request.Form.Get(url)
	if !strings.EqualFold(m, "") && !strings.EqualFold(u, "") {
		switch {
		case strings.EqualFold("GET", m) && strings.HasPrefix(u, "http"):
		case strings.EqualFold("PUT", m) && strings.HasPrefix(u, "http"):
		case strings.EqualFold("POST", m) && strings.HasPrefix(u, "http"):
		case strings.EqualFold("DELETE", m) && strings.HasPrefix(u, "http"):
		case strings.EqualFold("GET", m) && strings.HasPrefix(u, "https"):
		case strings.EqualFold("PUT", m) && strings.HasPrefix(u, "https"):
		case strings.EqualFold("POST", m) && strings.HasPrefix(u, "https"):
		case strings.EqualFold("DELETE", m) && strings.HasPrefix(u, "https"):
		case strings.EqualFold("GET", m) && strings.HasPrefix(u, "ftp"):
		case strings.EqualFold("PUT", m) && strings.HasPrefix(u, "ftp"):
		case strings.EqualFold("GET", m) && strings.HasPrefix(u, "ftps"):
		case strings.EqualFold("PUT", m) && strings.HasPrefix(u, "ftps"):
		default:
			log.Fatalf("accept wrong request %v\n only post was accepted", request)
			fmt.Fprintf(writer, "{\"code\":\"99\",\"desc\":\"wrong request %v\"\n", request)

		}
		return
	}
	log.Fatalf("accept wrong request %v\n only post was accepted", request)
	fmt.Fprintf(writer, "{\"code\":\"99\",\"desc\":\"wrong request %v\"\n", request)
	return
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

//func postOnlyFilter(writer http.ResponseWriter, r *http.Request) {
//
//}
//
//func FwrongRequestf(writer http.ResponseWriter, r *http.Request) {
//
//}
