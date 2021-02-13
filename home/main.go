package home

import (
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {

	// /srv/src/home
	//   Files:
	//     README.md
	//     main.go
	//     templates
	// /srv/src/serverless_function_app
	//   Files:
	//     main

	w.Header().Set("Content-Type", "text/html")
	home := "/srv/src/home/templates/index.html"
	http.ServeFile(w, r, home)
}
