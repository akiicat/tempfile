package upload

import (
  "os"
  "fmt"
  "log"
  "time"

  "net/http"

  "math/rand"
  "encoding/json"

  "github.com/akiirobot/tempfile/sign"
  "github.com/akiirobot/tempfile/db"
)


type Sign struct {
  Url    string `json:"url"`
  Token  string `json:"token"`
}


func UploadUrl(w http.ResponseWriter, r *http.Request) {
  if r.Method != "POST" {
    fmt.Printf("status method %s is not allowd, only allowed post", r.Method)
    http.Error(w, "405 - Status Method Not Allowed - " + r.Method, http.StatusMethodNotAllowed)
    return
  }

  objectName := r.PostFormValue("filename")
  token := randString(3)
  filename := token + "/" + objectName

  if filename == "" {
    log.Printf("Filename is empty")
    http.Error(w, "400 - Status Bad Request", http.StatusBadRequest)
    return
  }

  c, err := db.NewDatastoreDB(os.Getenv("PROJECT_ID"))
  if err != nil {
    http.Error(w, "500 - Status Internal Server Error", http.StatusInternalServerError)
    log.Printf("New Datastore DB by Project %s error: %v", os.Getenv("PROJECT_ID"), err)
    return
  }

  err = r.ParseForm()
  if err != nil {
    http.Error(w, "400 - Status Bad Request", http.StatusBadRequest)
    log.Printf("Request Parse Form error: %v", err)
    return
  }

  f := &db.File{
    FileName:   filename,
    Token:      token,
    CreatedAt:  time.Now(),
    UpdatedAt:  time.Now(),
  }

  k, err := c.AddFile(f)
  if err != nil {
    log.Printf("Datastore DB Add file %s/%s and Key is %s, error: %v", token, filename, k, err)
    http.Error(w, "500 - Status Internal Server Error - " + filename, http.StatusInternalServerError)
    return
  }

  url, err := sign.Sign(os.Getenv("SERVICE_JSON_FILE"), os.Getenv("BUCKET_NAME"), filename, "POST", 15)
  if err != nil {
    log.Fatalln(err)
    http.Error(w, "403 - Status Forbidden - " + err.Error(), http.StatusForbidden)
    return
  }

  obj := Sign{url, token}
  rtn, err := json.Marshal(obj)
  if err != nil {
		// return nil, fmt.Errorf("Json marshal %s, error: %v", fullname, err)
    log.Fatalln(err)
    http.Error(w, "500 - Status Forbidden - " + err.Error(), http.StatusInternalServerError)
    return
  }

  w.Header().Set("Content-Type", "application/json")
  w.Write(rtn)
}

func randString(n int) string {
  rand.Seed(time.Now().UTC().UnixNano())
  letterBytes := "abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNOPQRSTUVWXYZ"
  b := make([]byte, n)
  for i := range b {
    b[i] = letterBytes[rand.Intn(len(letterBytes))]
  }
  return string(b)
}



