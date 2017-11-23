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
	var m u string
	
	u = "http://localhost:8001/test?a=1&b=2"
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
	fmt.Println(resp.Status=="200 OK")//check the response status
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
	fmt.Println(resp.Status=="200 OK")//check the response status
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
	fmt.Println(resp.Status=="200 OK")//check the response status
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
	fmt.Println(resp.Status=="200 OK")//check the response status
	fmt.Printf("DELETE:\n %v\n", resp)
	buf = new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	s = buf.String()
	fmt.Println(s)
	buf.Reset()

	m = base64.StdEncoding.EncodeToString([]byte(http.MethodGet))
	u = "ftp://cme:CMEpassword1&@172.16.4.14:2121/a/readme.txt"
	u = base64.StdEncoding.EncodeToString([]byte(u))
	req, err = http.NewRequest(http.MethodPost, "http://localhost:8000/agent?m="+m+"&u="+u, nil)
	if nil != err {
		fmt.Printf("%v\n", err)
	}
	resp, err = client.Do(req)
	if nil != err {
		fmt.Printf("%v\n", err)
	}
	fmt.Println(resp.Status=="200 OK")//check the response status
	fmt.Printf("GET:\n %v\n", resp)
	buf = new(bytes.Buffer)
	n, err := buf.ReadFrom(resp.Body)
	if nil != err {
		fmt.Printf("%v\n", err)
	}
	fmt.Println(n)
	s = buf.String()
	fmt.Println(s)
	buf.Reset()

	m = base64.StdEncoding.EncodeToString([]byte(http.MethodDelete))
	u = "ftp://cme:CMEpassword1&@172.16.4.14:2121/a/readme.txt"
	u = base64.StdEncoding.EncodeToString([]byte(u))
	req, err = http.NewRequest(http.MethodPost, "http://localhost:8000/agent?m="+m+"&u="+u, nil)
	if nil != err {
		fmt.Printf("%v\n", err)
	}
	resp, err = client.Do(req)
	if nil != err {
		fmt.Printf("%v\n", err)
	}
	fmt.Println(resp.Status=="200 OK")//check the response status
	fmt.Printf("GET:\n %v\n", resp)
	buf = new(bytes.Buffer)
	n, err := buf.ReadFrom(resp.Body)
	if nil != err {
		fmt.Printf("%v\n", err)
	}
	fmt.Println(n)
	s = buf.String()
	fmt.Println(s)
	buf.Reset()
	
	
	m = base64.StdEncoding.EncodeToString([]byte(http.MethodPost))
	u = "ftp://cme:CMEpassword1&@172.16.4.14:2121/a/"
	u = base64.StdEncoding.EncodeToString([]byte(u))
	
	// Create buffer
    buf = new(bytes.Buffer) // caveat IMO dont use this for large files, \
    // create a tmpfile and assemble your multipart from there (not tested)
    w := multipart.NewWriter(buf)
    // Create file field
    fw, err := w.CreateFormFile("file", "helloworld.go") //这里的file很重要，必须和服务器端的FormFile一致
    if err != nil {
    fmt.Println("c")
    return err
    }
    fd, err := os.Open("helloworld.go")
    if err != nil {
    fmt.Println("d")
    return err
    }
    defer fd.Close()
    // Write file field from file to upload
    _, err = io.Copy(fw, fd)
    if err != nil {
    fmt.Println("e")
    return err
    }
    // Important if you do not close the multipart writer you will not have a
    // terminating boundry
    w.Close()
    //req, err := http.NewRequest("POST","http://192.168.2.127/configure.go?portId=2", buf)
    req, err = http.NewRequest(http.MethodPost, "http://localhost:8000/agent?m="+m+"&u="+u, buf)
    if nil != err {
		fmt.Printf("%v\n", err)
	}
    req.Header.Set("Content-Type", w.FormDataContentType())
    resp, err := client.Do(req)
	if nil != err {
		fmt.Printf("%v\n", err)
	}
	fmt.Println(resp.Status=="200 OK")//check the response status
	fmt.Printf("POST:\n %v\n", resp)
	buf = new(bytes.Buffer)
	n, err := buf.ReadFrom(resp.Body)
	if nil != err {
		fmt.Printf("%v\n", err)
	}
	fmt.Println(n)
	s = buf.String()
	fmt.Println(s)
	buf.Reset()
	
}
