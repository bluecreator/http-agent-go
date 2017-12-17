# agent-go

## What's is agent-go?
agent-go is an open source agent(proxy and adapter) that written in golang.

## Features    
- [X] READY     
- [ ] TODO 
  
### agent features
- [X] HTTP proxy(support GET PUT POST DELETE to target server)    
- [ ] HTTPS proxy(support GET PUT POST DELETE to target server)    
- [X] FTP adaptor(support GET POST DELETE to target server)    
- [ ] SFTP adaptor(support GET POST DELETE to target server)    

### admin features
- [ ] black target list(host:port)    
- [ ] white target list(host:port)    
- [ ] access list(host)    
- [ ] target key management   
- [ ] download log file  
- [ ] client and target mapping(host->host:port)
 
## Example     
First, put target method(m), target url(u) to agent server's url as query string.    
Then, write the body follow the target server's requirements.    
Finally, get the response from target server.    

```go
	var m, u string

	u = "http://localhost:8001/test?a=1&b=2"
	m = base64.StdEncoding.EncodeToString([]byte("PUT"))
	u = base64.StdEncoding.EncodeToString([]byte(u))
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, "http://localhost:8000/agent?m="+m+"&u="+u, strings.NewReader("PUT http://localhost:8001/test?a=1&b=2"))
	if nil != err {
		fmt.Printf("%v\n", err)
		return
	}
	resp, err := client.Do(req)
	if nil != err {
		fmt.Printf("%v\n", err)
		return
	}
	fmt.Println(resp.StatusCode == 200) //check the response status
	fmt.Printf("PUT:\n %v\n", resp)
	buf := new(bytes.Buffer)
	n, err := buf.ReadFrom(resp.Body)
	if nil != err {
		fmt.Printf("%v\n", err)
		return
	}
	fmt.Println(n)
	s := buf.String()
	fmt.Println(s)
	buf.Reset()
	
	
	m = base64.StdEncoding.EncodeToString([]byte(http.MethodGet))
	u = "ftp://cme:CMEpassword1&@172.16.4.14:2121/a/admin.go"
	u = base64.StdEncoding.EncodeToString([]byte(u))
	req, err = http.NewRequest(http.MethodPost, "http://localhost:8000/agent?m="+m+"&u="+u, nil)
	if nil != err {
		fmt.Printf("%v\n", err)
		return
	}
	resp, err = client.Do(req)
	if nil != err {
		fmt.Printf("%v\n", err)
		return
	}
	fmt.Println(resp.StatusCode == 200) //check the response status
	fmt.Printf("GET:\n %v\n", resp)
	buf = new(bytes.Buffer)
	n, err = buf.ReadFrom(resp.Body)
	if nil != err {
		fmt.Printf("%v\n", err)
		return
	}
	fmt.Println(n)
	s = buf.String()
	fmt.Println(s)
	buf.Reset()
```

