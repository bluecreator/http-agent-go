package agent

import (
	"fmt"
	"net/http"
)

func admin(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Host, r.Method, r.URL.Path)
	r.ParseForm()
	fmt.Fprintf(w, "%v\n", r.Form)
}
