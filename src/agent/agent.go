package agent

import (
	"encoding/base64"
	"fmt"
	"github.com/jlaffaye/ftp"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

const (
	METHOD = "m"
	URL    = "u"
)

func agent(w http.ResponseWriter, r *http.Request) {

	if http.MethodPost != r.Method {
		log.Fatalf("accept wrong request %v\n only post was accepted", r)
		error := fmt.Sprintf("Only POST method was accepted but %s reveived", r.Method)
		http.Error(w, error, http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); nil != err {
		log.Fatalf("%v\n", err)
		error := fmt.Sprintf("%v\n", err)
		http.Error(w, error, http.StatusBadRequest)
		return
	}
	//m := r.Form.Get(METHOD)
	//u := r.Form.Get(URL)
	mbytes, err := base64.StdEncoding.DecodeString(r.Form.Get(METHOD))
	if nil != err {
		log.Fatalf("%v\n", err)
		error := fmt.Sprintf("%v\n", err)
		http.Error(w, error, http.StatusBadRequest)
		return
	}
	ubytes, err := base64.StdEncoding.DecodeString(r.Form.Get(URL))
	if nil != err {
		log.Fatalf("%v\n", err)
		error := fmt.Sprintf("%v\n", err)
		http.Error(w, error, http.StatusBadRequest)
		return
	}
	m := string(mbytes)
	u := string(ubytes)
	if !strings.EqualFold(m, "") && !strings.EqualFold(u, "") {
		switch {
		case (strings.EqualFold(http.MethodGet, m) || strings.EqualFold(http.MethodDelete, m)) && strings.HasPrefix(u, "http"):
			req, err := http.NewRequest(m, u, nil)
			if nil != err {
				log.Fatalf("%v\n", err)
				error := fmt.Sprintf("%v\n", err)
				http.Error(w, error, http.StatusBadRequest)
				return
			}
			client := &http.Client{}
			resp, err := client.Do(req)
			if nil != err {
				log.Fatalf("%v\n", err)
				error := fmt.Sprintf("%v\n", err)
				http.Error(w, error, http.StatusBadRequest)
				return
			}
			_, err = io.Copy(w, resp.Body)
			defer resp.Body.Close()
			if nil != err {
				log.Fatalf("%v\n", err)
				error := fmt.Sprintf("%v\n", err)
				http.Error(w, error, http.StatusBadRequest)
			}
		case (strings.EqualFold(http.MethodPut, m) || strings.EqualFold(http.MethodPost, m)) && strings.HasPrefix(u, "http"):
			req, err := http.NewRequest(m, u, r.Body)
			if nil != err {
				log.Fatalf("%v\n", err)
				error := fmt.Sprintf("%v\n", err)
				http.Error(w, error, http.StatusBadRequest)
				return
			}
			client := &http.Client{}
			resp, err := client.Do(req)
			if nil != err {
				log.Fatalf("%v\n", err)
				error := fmt.Sprintf("%v\n", err)
				http.Error(w, error, http.StatusBadRequest)
				return
			}
			_, err = io.Copy(w, resp.Body)
			defer resp.Body.Close()
			if nil != err {
				log.Fatalf("%v\n", err)
				error := fmt.Sprintf("%v\n", err)
				http.Error(w, error, http.StatusBadRequest)
				return
			}
		case strings.EqualFold(http.MethodGet, m) && strings.HasPrefix(u, "https"):
		case strings.EqualFold(http.MethodPut, m) && strings.HasPrefix(u, "https"):
		case strings.EqualFold(http.MethodPost, m) && strings.HasPrefix(u, "https"):
		case strings.EqualFold(http.MethodDelete, m) && strings.HasPrefix(u, "https"):
		case strings.EqualFold(http.MethodGet, m) && strings.HasPrefix(u, "ftp"):
			ftpUrl, err := url.Parse(u)
			if nil != err {
				log.Fatalf("%v\n", err)
				error := fmt.Sprintf("%v\n", err)
				http.Error(w, error, http.StatusBadRequest)
				return
			}
			log.Fatalf("%s %s %v\n", ftpUrl.Host, ftpUrl.Path, ftpUrl.User)
			server, err := ftp.Connect(ftpUrl.Host)
			if nil != err {
				log.Fatalf("%v\n", err)
				error := fmt.Sprintf("%v\n", err)
				http.Error(w, error, http.StatusBadRequest)
				return
			}
			user := ftpUrl.User.Username()
			if strings.EqualFold(user, "") {
				user = "anonymous"
			}
			pass, _ := ftpUrl.User.Password()
			err = server.Login(user, pass)
			if nil != err {
				log.Fatalf("%v\n", err)
				error := fmt.Sprintf("%v\n", err)
				http.Error(w, error, http.StatusBadRequest)
				return
			}

			resp, err := server.Retr(ftpUrl.Path)
			if nil != err {
				log.Fatalf("%v\n", err)
				error := fmt.Sprintf("%v\n", err)
				http.Error(w, error, http.StatusBadRequest)
				return
			}
			_, err = io.Copy(w, resp)
			defer resp.Close()
			if nil != err {
				log.Fatalf("%v\n", err)
				error := fmt.Sprintf("%v\n", err)
				http.Error(w, error, http.StatusBadRequest)
				return
			}

		case strings.EqualFold(http.MethodPut, m) && strings.HasPrefix(u, "ftp"):
			ftpUrl, err := url.Parse(u)
			if nil != err {
				log.Fatalf("%v\n", err)
				error := fmt.Sprintf("%v\n", err)
				http.Error(w, error, http.StatusBadRequest)
				return
			}
			log.Fatalf("%s %s %v\n", ftpUrl.Host, ftpUrl.Path, ftpUrl.User)
			server, err := ftp.Connect(ftpUrl.Host)
			if nil != err {
				log.Fatalf("%v\n", err)
				error := fmt.Sprintf("%v\n", err)
				http.Error(w, error, http.StatusBadRequest)
				return
			}
			user := ftpUrl.User.Username()
			if strings.EqualFold(user, "") {
				user = "anonymous"
			}
			pass, _ := ftpUrl.User.Password()
			err = server.Login(user, pass)
			if nil != err {
				log.Fatalf("%v\n", err)
				error := fmt.Sprintf("%v\n", err)
				http.Error(w, error, http.StatusBadRequest)
				return
			}

			resp, err := server.Retr(ftpUrl.Path)
			if nil != err {
				log.Fatalf("%v\n", err)
				error := fmt.Sprintf("%v\n", err)
				http.Error(w, error, http.StatusBadRequest)
				return
			}
			_, err = io.Copy(w, resp)
			defer resp.Close()
			if nil != err {
				log.Fatalf("%v\n", err)
				error := fmt.Sprintf("%v\n", err)
				http.Error(w, error, http.StatusBadRequest)
				return
			}
			err = server.Logout()
			if nil != err {
				log.Fatalf("%v\n", err)
				error := fmt.Sprintf("%v\n", err)
				http.Error(w, error, http.StatusBadRequest)
				return
			}
			err = server.Quit()
			if nil != err {
				log.Fatalf("%v\n", err)
				error := fmt.Sprintf("%v\n", err)
				http.Error(w, error, http.StatusBadRequest)
				return
			}

		case strings.EqualFold(http.MethodDelete, m) && strings.HasPrefix(u, "ftp"):
		case strings.EqualFold(http.MethodGet, m) && strings.HasPrefix(u, "ftps"):
		case strings.EqualFold(http.MethodPut, m) && strings.HasPrefix(u, "ftps"):
		case strings.EqualFold(http.MethodDelete, m) && strings.HasPrefix(u, "ftps"):
		default:
			log.Fatalf("accept wrong request %v\n only post was accepted", r)
			fmt.Fprintf(w, "{\"code\":\"99\",\"desc\":\"wrong request %v\"\n", r)
		}
	} else {
		log.Fatalf("accept wrong request %v\n only post was accepted", r)
		fmt.Fprintf(w, "{\"code\":\"99\",\"desc\":\"wrong request %v\"\n", r)
	}
	return
	//	ftp, err := ftp.Connect("ftp://192.168.1.1:21")
	//	if nil != err {
	//		log.Fatalf("%v\n", err)
	//		fmt.Fprintf(w, "%v\n", err)
	//	}
	//	ftp.RetrFrom(" ", 0)
}

//func forward(m, u string, w http.ResponseWriter, body io.ReadCloser) {
//	resp, err := http.NewRequest(m, u, nil)
//	if nil != err {
//		log.Fatalf("%v\n", err)
//		error := fmt.Sprintf("%v\n", err)
//		http.Error(w, error, StatusBadRequest)
//		return
//	}
//	client := &http.Client{}
//	resp, err := client.Do(req)
//	if nil != err {
//		log.Fatalf("%v\n", err)
//		error := fmt.Sprintf("%v\n", err)
//		http.Error(w, error, StatusBadRequest)
//		return
//	}
//	_, err = io.Copy(w, resp.Body)
//	if nil != err {
//		log.Fatalf("%v\n", err)
//		error := fmt.Sprintf("%v\n", err)
//		http.Error(w, error, StatusBadRequest)
//	}
//}
//
//func adapt(m, u string, w io.Writer, r io.Reader) {
//
//}
