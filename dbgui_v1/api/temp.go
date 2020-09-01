package api

import (
	"net/http"
)

func Render(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/page-login.html")

}
