# agent-go

## What's is agent-go?
agent-go is an open source agent(proxy and adapter) that written in golang.

## Features    
- [X] READY     
- [ ] TODO 
  
### agent features
- [X] HTTP proxy ( support GET PUT POST DELETE)    
- [ ] HTTPS proxy ( support GET PUT POST DELETE)    
- [X] FTP adaptor ( support GET POST DELETE)    
- [ ] FTP adaptor ( support GET POST DELETE)    
- [ ] FTP adaptor ( support GET POST DELETE)    

### admin features
- [ ] black target list ( scheme://host:port)    
- [ ] white target list ( scheme://host:port)    
- [ ] access list ( method:host)    
- [ ] target key management   
 
## Example     
First, put target method(m), target url(u) to agent server's url as query string.    
Then, write the body follow the target server's requirements.    
Finally, get the response from target server.    

```go
	var m, u string

	u = "http://localhost:8001/test?a=1&b=2"
	m = base64.StdEncoding.EncodeToString([]byte(http.MethodPut))
	u = base64.StdEncoding.EncodeToString([]byte(u))
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, "http://localhost:8000/agent?m="+m+"&u="+u, strings.NewReader("methodPut"))
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
```

