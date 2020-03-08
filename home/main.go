package home

import (
  "net/http"
)


func HomePage(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "text/html")
  home := "/srv/vendor/home/templates/index.html"
  http.ServeFile(w, r, home)
}

