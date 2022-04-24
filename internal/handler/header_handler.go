package handler

import "net/http"

func HeaderHandler(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		w.Header()[name] = headers
	}
}
