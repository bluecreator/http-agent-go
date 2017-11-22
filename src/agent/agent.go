package agent

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

const (
	method = "method"
	url    = "url"
	post   = "POST"
	get    = "GET"
	put    = "PUT"
	delete = "DELETE"
)

func agent(writer http.ResponseWriter, request *http.Request) {

	if post != request.Method {
		log.Fatalf("accept wrong request %v\n only post was accepted", request)
		fmt.Fprintf(writer, "{\"code\":\"99\",\"desc\":\"wrong request %v\"\n", request)
	} else {
		err := request.ParseForm()
		if nil != err {
			log.Fatalf("%v\n", err)
			fmt.Fprintf(writer, "{\"code\":\"99\",\"desc\":\"%v\"\n", err)
		} else {
			m := request.Form.Get(method)
			u := request.Form.Get(url)
			if !strings.EqualFold(m, "") && !strings.EqualFold(u, "") {
				switch {
				case (strings.EqualFold("GET", m) || strings.EqualFold("DELETE", m)) && strings.HasPrefix(u, "http"):
					resp, err := http.NewRequest(m, u, nil)
					if nil != err {
						log.Fatalf("%v\n", err)
						fmt.Fprintf(writer, "{\"code\":\"99\",\"desc\":\"%v\"\n", err)
					} else {
						_, err = io.Copy(writer, resp.Body)
						if nil != err {
							log.Fatalf("%v\n", err)
							fmt.Fprintf(writer, "{\"code\":\"99\",\"desc\":\"%v\"\n", err)
						}
					}

				case (strings.EqualFold("PUT", m) || strings.EqualFold("POST", m)) && strings.HasPrefix(u, "http"):
					resp, err := http.NewRequest(m, u, request.Body)
					if nil != err {
						log.Fatalf("%v\n", err)
						fmt.Fprintf(writer, "{\"code\":\"99\",\"desc\":\"%v\"\n", err)
					} else {
						_, err = io.Copy(writer, resp.Body)
						if nil != err {
							log.Fatalf("%v\n", err)
							fmt.Fprintf(writer, "{\"code\":\"99\",\"desc\":\"%v\"\n", err)
						}
					}
				case strings.EqualFold("GET", m) && strings.HasPrefix(u, "https"):
				case strings.EqualFold("PUT", m) && strings.HasPrefix(u, "https"):
				case strings.EqualFold("POST", m) && strings.HasPrefix(u, "https"):
				case strings.EqualFold("DELETE", m) && strings.HasPrefix(u, "https"):
				case strings.EqualFold("GET", m) && strings.HasPrefix(u, "ftp"):
				case strings.EqualFold("PUT", m) && strings.HasPrefix(u, "ftp"):
				case strings.EqualFold("DELETE", m) && strings.HasPrefix(u, "ftp"):
				case strings.EqualFold("GET", m) && strings.HasPrefix(u, "ftps"):
				case strings.EqualFold("PUT", m) && strings.HasPrefix(u, "ftps"):
				case strings.EqualFold("DELETE", m) && strings.HasPrefix(u, "ftps"):
				default:
					log.Fatalf("accept wrong request %v\n only post was accepted", request)
					fmt.Fprintf(writer, "{\"code\":\"99\",\"desc\":\"wrong request %v\"\n", request)

				}
			} else {
				log.Fatalf("accept wrong request %v\n only post was accepted", request)
				fmt.Fprintf(writer, "{\"code\":\"99\",\"desc\":\"wrong request %v\"\n", request)
			}
		}
	}
	return
	//	ftp, err := ftp.Connect("ftp://192.168.1.1:21")
	//	if nil != err {
	//		log.Fatalf("%v\n", err)
	//		fmt.Fprintf(writer, "%v\n", err)
	//	}
	//	ftp.RetrFrom(" ", 0)
}
