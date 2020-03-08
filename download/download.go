package download

import (
  "os"
  "log"
  "fmt"

  "net/http"
  "net/url"

  "github.com/akiirobot/tempfile/db"
)

func DownloadUrl(w http.ResponseWriter, r *http.Request) {
  if r.Method != "GET" {
    log.Printf("status method %s is not allowd, only allowed get", r.Method)
    http.Error(w, "405 - Status Method Not Allowed - " + r.Method, http.StatusMethodNotAllowed)
    return
  }

  project := os.Getenv("PROJECT_ID")

  c, err := db.NewDatastoreDB(project)
  if err != nil {
    log.Printf("New Datastore DB error: %v", err)
    http.Error(w, "500 - Status Internal Server Error", http.StatusInternalServerError)
    return
  }

  token := r.URL.Path[1:]

  file, err := c.GetFileByToken(token)
  if err != nil {
    log.Printf("400 - Status Bad Request: %v", err)
    http.Error(w, "400 - Status Bad Request", http.StatusBadRequest)
    return
  }

  link := sign(file.Token, file.FileName)

  // cors
  cs := w.Header().Get("Set-Cookie")
  cs += "; SameSite=None; Secure"
  w.Header().Set("Set-Cookie", cs)

  log.Printf("%s Redirect to: %s -> %s", file.FileName, token, link)
  http.Redirect(w, r, link, http.StatusFound)
}

func sign(token, file string) string {
  return fmt.Sprintf("https://storage.googleapis.com/%s/%s/%s", os.Getenv("BUCKET_NAME"), token, url.PathEscape(file))
}

