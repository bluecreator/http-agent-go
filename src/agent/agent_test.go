package agent

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"
	"testing"
)

func Test_agent(t *testing.T) {
	var m string
	var u = "http://localhost:8001/test?a=1&b=2"
	m = base64.StdEncoding.EncodeToString([]byte(http.MethodPut))
	u = base64.StdEncoding.EncodeToString([]byte(u))
	fmt.Println(m)
	fmt.Println(u)
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, "http://localhost:8000/agent?m="+m+"&u="+u, strings.NewReader("methodPut"))
	if nil != err {
		fmt.Printf("%v\n", err)
	}
	resp, err := client.Do(req)
	if nil != err {
		fmt.Printf("%v\n", err)
	}
	fmt.Printf("PUT:\n %v\n", resp)
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	s := buf.String()
	fmt.Println(s)
	buf.Reset()
	m = base64.StdEncoding.EncodeToString([]byte(http.MethodPost))
	req, err = http.NewRequest(http.MethodPost, "http://localhost:8000/agent?m="+m+"&u="+u, strings.NewReader("methodPost"))
	if nil != err {
		fmt.Printf("%v\n", err)
	}
	resp, err = client.Do(req)
	if nil != err {
		fmt.Printf("%v\n", err)
	}
	fmt.Printf("POST:\n %v\n", resp)
	buf = new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	s = buf.String()
	fmt.Println(s)
	buf.Reset()
	m = base64.StdEncoding.EncodeToString([]byte(http.MethodGet))
	req, err = http.NewRequest(http.MethodPost, "http://localhost:8000/agent?m="+m+"&u="+u, nil)
	if nil != err {
		fmt.Printf("%v\n", err)
	}
	resp, err = client.Do(req)
	if nil != err {
		fmt.Printf("%v\n", err)
	}
	fmt.Printf("GET:\n %v\n", resp)
	buf = new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	s = buf.String()
	fmt.Println(s)
	buf.Reset()
	m = base64.StdEncoding.EncodeToString([]byte(http.MethodDelete))
	req, err = http.NewRequest(http.MethodPost, "http://localhost:8000/agent?m="+m+"&u="+u, nil)
	if nil != err {
		fmt.Printf("%v\n", err)
	}
	resp, err = client.Do(req)
	if nil != err {
		fmt.Printf("%v\n", err)
	}
	fmt.Printf("DELETE:\n %v\n", resp)
	buf = new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	s = buf.String()
	fmt.Println(s)
	buf.Reset()
}
