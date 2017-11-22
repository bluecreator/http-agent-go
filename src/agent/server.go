package agent

import (
	"log"
	"net/http"
)

func StartServer() {
	http.HandleFunc("/agent", agent)
	http.HandleFunc("/admin", admin)
	log.Fatalln(http.ListenAndServe("172.16.4.14:8000", nil))
}
