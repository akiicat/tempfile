package upload

import (
	"log"
	"os"
	"time"

	"net/http"
	"net/url"

	"encoding/json"
	"math/rand"

	"github.com/akiirobot/tempfile/db"
)

type Sign struct {
	Url   string `json:"url"`
	Token string `json:"token"`
	Name  string `json:"name"`
	Nonce string `json:"nonce"`
}

func UploadUrl(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		log.Printf("status method %s is not allowd, only allowed post", r.Method)
		http.Error(w, "405 - Status Method Not Allowed - "+r.Method, http.StatusMethodNotAllowed)
		return
	}

	project := os.Getenv("PROJECT_ID")
	bucket := os.Getenv("BUCKET_NAME")

	filename := r.PostFormValue("filename")
	filesize := r.PostFormValue("filesize")
	token := randString(3)
	fullname := token + "/" + filename
	nonce := r.PostFormValue("nonce")

	if fullname == "" {
		log.Printf("Filename is empty")
		http.Error(w, "400 - Status Bad Request", http.StatusBadRequest)
		return
	}

	c, err := db.NewDatastoreDB(project)
	if err != nil {
		http.Error(w, "500 - Status Internal Server Error", http.StatusInternalServerError)
		log.Printf("New Datastore DB by Project %s error: %v", project, err)
		return
	}

	err = r.ParseForm()
	if err != nil {
		http.Error(w, "400 - Status Bad Request", http.StatusBadRequest)
		log.Printf("Request Parse Form error: %v", err)
		return
	}

	f := &db.File{
		FileName:  filename,
		FileSize:  filesize,
		Token:     token,
		CreatedAt: time.Now(),
	}

	k, err := c.AddFile(f)
	if err != nil {
		log.Printf("Datastore DB Add file %s/%s and Key is %s, error: %v", token, filename, k, err)
		http.Error(w, "500 - Status Internal Server Error - "+filename, http.StatusInternalServerError)
		return
	}

	obj := newSign(bucket, token, filename, nonce)
	rtn, err := json.Marshal(obj)
	if err != nil {
		log.Fatalln("JSON marshal error: ", err)
		http.Error(w, "500 - Status Forbidden - "+err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("response", rtn)
	w.Header().Set("Content-Type", "application/json")
	w.Write(rtn)
}

func newSign(bucket, token, filename, nonce string) Sign {
	return Sign{
		signPublic(bucket, token, filename),
		token,
		url.PathEscape(filename),
		nonce,
	}
}

func signPublic(bucket, token, name string) string {
	return "https://storage.googleapis.com/" + bucket + "/" + token + "/" + url.PathEscape(name)
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
