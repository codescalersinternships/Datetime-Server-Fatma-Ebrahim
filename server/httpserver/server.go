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
		datetime := time.Now().Local().Format("Monday 02-01-2006 15:04:05")
		data, _ := json.Marshal(map[string]string{"datetime": datetime})
		fmt.Fprint(w, string(data))
	} else {

		fmt.Fprint(w, time.Now().Local().Format("Monday 02-01-2006 15:04:05"))
	}
}

// func main() {
// 	log.Fatal(http.ListenAndServe(":8000", http.HandlerFunc(HTTPHandler)))

// }
