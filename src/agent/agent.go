package agent

import (
	"encoding/base64"
	"fmt"
	"github.com/jlaffaye/ftp"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const (
	METHOD = "m"
	URL    = "u"
)

func agent(w http.ResponseWriter, r *http.Request) {

	if http.MethodPost != r.Method {
		log.Printf("accept wrong request %v\n only post was accepted", r)
		error := fmt.Sprintf("Only POST method was accepted but %s reveived", r.Method)
		http.Error(w, error, http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); nil != err {
		log.Printf("%v\n", err)
		error := fmt.Sprintf("%v\n", err)
		http.Error(w, error, http.StatusBadRequest)
		return
	}
	//m := r.Form.Get(METHOD)
	//u := r.Form.Get(URL)
	mbytes, err := base64.StdEncoding.DecodeString(r.Form.Get(METHOD))
	if nil != err {
		log.Printf("%v\n", err)
		error := fmt.Sprintf("%v\n", err)
		http.Error(w, error, http.StatusBadRequest)
		return
	}
	ubytes, err := base64.StdEncoding.DecodeString(r.Form.Get(URL))
	if nil != err {
		log.Printf("%v\n", err)
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
				log.Printf("%v\n", err)
				error := fmt.Sprintf("%v\n", err)
				http.Error(w, error, http.StatusBadRequest)
				return
			}
			client := &http.Client{}
			resp, err := client.Do(req)
			if nil != err {
				log.Printf("%v\n", err)
				error := fmt.Sprintf("%v\n", err)
				http.Error(w, error, http.StatusBadRequest)
				return
			}
			_, err = io.Copy(w, resp.Body)
			defer resp.Body.Close()
			if nil != err {
				log.Printf("%v\n", err)
				error := fmt.Sprintf("%v\n", err)
				http.Error(w, error, http.StatusBadRequest)
			}
		case (strings.EqualFold(http.MethodPut, m) || strings.EqualFold(http.MethodPost, m)) && strings.HasPrefix(u, "http"):
			req, err := http.NewRequest(m, u, r.Body)
			if nil != err {
				log.Printf("%v\n", err)
				error := fmt.Sprintf("%v\n", err)
				http.Error(w, error, http.StatusBadRequest)
				return
			}
			client := &http.Client{}
			resp, err := client.Do(req)
			if nil != err {
				log.Printf("%v\n", err)
				error := fmt.Sprintf("%v\n", err)
				http.Error(w, error, http.StatusBadRequest)
				return
			}
			_, err = io.Copy(w, resp.Body)
			defer resp.Body.Close()
			if nil != err {
				log.Printf("%v\n", err)
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
				log.Printf("%v\n", err)
				error := fmt.Sprintf("%v\n", err)
				http.Error(w, error, http.StatusBadRequest)
				return
			}
			log.Printf("%s %s %v\n", ftpUrl.Host, ftpUrl.Path, ftpUrl.User)
			server, err := ftp.Connect(ftpUrl.Host)
			defer server.Quit()
			if nil != err {
				log.Printf("%v\n", err)
				error := fmt.Sprintf("%v\n", err)
				http.Error(w, error, http.StatusBadRequest)
				return
			}
			var user, pass string
			if nil == ftpUrl.User {
				user = "anonymous"
				pass = ""
			} else {
				user = ftpUrl.User.Username()
				if strings.EqualFold(user, "") {
					user = "anonymous"
				}
				pass, _ = ftpUrl.User.Password()
			}

			err = server.Login(user, pass)
			defer server.Logout()
			if nil != err {
				log.Printf("%v\n", err)
				error := fmt.Sprintf("%v\n", err)
				http.Error(w, error, http.StatusBadRequest)
				return
			}

			resp, err := server.Retr(ftpUrl.Path)
			if nil != err {
				log.Printf("%v\n", err)
				error := fmt.Sprintf("%v\n", err)
				http.Error(w, error, http.StatusBadRequest)
				return
			}
			_, err = io.Copy(w, resp)
			defer resp.Close()
			if nil != err {
				log.Printf("%v\n", err)
				error := fmt.Sprintf("%v\n", err)
				http.Error(w, error, http.StatusBadRequest)
				return
			}

		case strings.EqualFold(http.MethodPost, m) && strings.HasPrefix(u, "ftp"):
			err = r.ParseMultipartForm(16 << 10) //16M
			if nil != err {
				log.Printf("%v\n", err)
				error := fmt.Sprintf("%v\n", err)
				http.Error(w, error, http.StatusBadRequest)
				return
			}
			if nil == r.MultipartForm {
				log.Println("Not a MultipartForm.\n")
				http.Error(w, "Not a MultipartForm.\n", http.StatusBadRequest)
				return
			}
			if nil == r.MultipartForm.File || len(r.MultipartForm.File) == 0 {
				log.Println("MultipartForm did not have file.\n")
				http.Error(w, "MultipartForm did not have file.\n", http.StatusBadRequest)
				return
			}

			ftpUrl, err := url.Parse(u)
			if nil != err {
				log.Printf("%v\n", err)
				error := fmt.Sprintf("%v\n", err)
				http.Error(w, error, http.StatusBadRequest)
				return
			}
			log.Printf("%s %s %v\n", ftpUrl.Host, ftpUrl.Path, ftpUrl.User)
			server, err := ftp.Connect(ftpUrl.Host)
			defer server.Quit()
			if nil != err {
				log.Printf("%v\n", err)
				error := fmt.Sprintf("%v\n", err)
				http.Error(w, error, http.StatusBadRequest)
				return
			}
			var user, pass string
			if nil == ftpUrl.User {
				user = "anonymous"
				pass = ""
			} else {
				user = ftpUrl.User.Username()
				if strings.EqualFold(user, "") {
					user = "anonymous"
				}
				pass, _ = ftpUrl.User.Password()
			}
			err = server.Login(user, pass)
			defer server.Logout()
			if nil != err {
				log.Printf("%v\n", err)
				error := fmt.Sprintf("%v\n", err)
				http.Error(w, error, http.StatusBadRequest)
				return
			}

			/*
				RFC7578
				For form data that represents the content of a file, a name for the
				file SHOULD be supplied as well, by using a "filename" parameter of
				the Content-Disposition header field.  The file name isn't mandatory
				for cases where the file name isn't available or is meaningless or
				private; this might result, for example, when selection or drag-and-
				drop is used or when the form data content is streamed directly from
				a device.
			*/
			for k, v := range r.MultipartForm.File {
				for i := 0; i < len(v); i++ {
					f, err := v[i].Open()
					if nil != err {
						log.Printf("%v\n", err)
						error := fmt.Sprintf("%v\n", err)
						http.Error(w, error, http.StatusBadRequest)
						return
					}
					var filepath = v[i].Filename
					if strings.EqualFold(filepath, "") {
						filepath = k + strconv.Itoa(i)
					}
					if !strings.EqualFold(ftpUrl.Path, "") {
						filepath = ftpUrl.Path + "/" + filepath
					}
					err = server.Stor(filepath, f)
					if nil != err {
						log.Printf("%v\n", err)
						error := fmt.Sprintf("%v\n", err)
						http.Error(w, error, http.StatusBadRequest)
						return
					}
				}
			}

		case strings.EqualFold(http.MethodDelete, m) && strings.HasPrefix(u, "ftp"):
			ftpUrl, err := url.Parse(u)
			if nil != err {
				log.Printf("%v\n", err)
				error := fmt.Sprintf("%v\n", err)
				http.Error(w, error, http.StatusBadRequest)
				return
			}
			//			log.Printf("%s %s %v\n", ftpUrl.Host, ftpUrl.Path, ftpUrl.User)
			server, err := ftp.Connect(ftpUrl.Host)
			defer server.Quit()
			if nil != err {
				log.Printf("%v\n", err)
				error := fmt.Sprintf("%v\n", err)
				http.Error(w, error, http.StatusBadRequest)
				return
			}
			var user, pass string
			if nil == ftpUrl.User {
				user = "anonymous"
				pass = ""
			} else {
				user = ftpUrl.User.Username()
				if strings.EqualFold(user, "") {
					user = "anonymous"
				}
				pass, _ = ftpUrl.User.Password()
			}
			err = server.Login(user, pass)
			defer server.Logout()
			if nil != err {
				log.Printf("%v\n", err)
				error := fmt.Sprintf("%v\n", err)
				http.Error(w, error, http.StatusBadRequest)
				return
			}
			err = server.Delete(ftpUrl.Path)
			if nil != err {
				log.Printf("%v\n", err)
				error := fmt.Sprintf("%v\n", err)
				http.Error(w, error, http.StatusBadRequest)
				return
			}
		case strings.EqualFold(http.MethodGet, m) && strings.HasPrefix(u, "ftps"):
		case strings.EqualFold(http.MethodPut, m) && strings.HasPrefix(u, "ftps"):
		case strings.EqualFold(http.MethodDelete, m) && strings.HasPrefix(u, "ftps"):
		default:
			log.Printf("%v\n", err)
			error := fmt.Sprintf("%v\n", err)
			http.Error(w, error, http.StatusBadRequest)
		}
	} else {
		log.Printf("%v\n", err)
		error := fmt.Sprintf("%v\n", err)
		http.Error(w, error, http.StatusBadRequest)
	}
	return
	//	ftp, err := ftp.Connect("ftp://192.168.1.1:21")
	//	if nil != err {
	//		log.Printf("%v\n", err)
	//		fmt.Fprintf(w, "%v\n", err)
	//	}
	//	ftp.RetrFrom(" ", 0)
}

//func forward(m, u string, w http.ResponseWriter, body io.ReadCloser) {
//	resp, err := http.NewRequest(m, u, nil)
//	if nil != err {
//		log.Printf("%v\n", err)
//		error := fmt.Sprintf("%v\n", err)
//		http.Error(w, error, StatusBadRequest)
//		return
//	}
//	client := &http.Client{}
//	resp, err := client.Do(req)
//	if nil != err {
//		log.Printf("%v\n", err)
//		error := fmt.Sprintf("%v\n", err)
//		http.Error(w, error, StatusBadRequest)
//		return
//	}
//	_, err = io.Copy(w, resp.Body)
//	if nil != err {
//		log.Printf("%v\n", err)
//		error := fmt.Sprintf("%v\n", err)
//		http.Error(w, error, StatusBadRequest)
//	}
//}
//
//func adapt(m, u string, w io.Writer, r io.Reader) {
//
//}
