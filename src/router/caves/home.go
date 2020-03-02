package caves

import "net/http"

// HomeHandler Function
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
