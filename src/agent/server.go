package agent

import ()

func StartServer() {
	http.HandleFunc("/agent", agent)
	http.HandleFunc("/admin", admin)
	log.Fatalln(http.ListenAndServe("localhost:8000", nil))
}
