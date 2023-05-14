package handler

import (
	"io"
	"net/http"
)

func Index(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "<h1>Welcome</h1>\n")
}
