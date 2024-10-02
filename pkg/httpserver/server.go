package httpserver

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func HTTPHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/datetime" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	contenttype := r.Header.Get("Content-Type")
	if contenttype == "application/json" {
		datetime := time.Now().Local().Format(time.RubyDate)
		data, _ := json.Marshal(map[string]string{"datetime": datetime})
		fmt.Fprint(w, string(data))
	} else {

		fmt.Fprint(w, time.Now().Local().Format(time.RubyDate))
	}
}
