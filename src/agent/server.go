package agent

import (
	"log"
	"net/http"
)

func StartServer() {
	http.HandleFunc("/agent", agent)
	http.HandleFunc("/admin", admin)
	log.Fatalln(http.ListenAndServe(":8000", nil))
}
