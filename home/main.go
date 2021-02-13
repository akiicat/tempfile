package home

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Fprintln(w, path)

	files, err := ioutil.ReadDir(".")
	if err != nil {
		http.Error(w, "Unable to read files", http.StatusInternalServerError)
		log.Printf("ioutil.ListFiles: %v", err)
		return
	}
	fmt.Fprintln(w, "Files:")
	for _, f := range files {
		fmt.Fprintf(w, "\t%v\n", f.Name())
	}

	// w.Header().Set("Content-Type", "text/html")
	// home := "/srv/vendor/home/templates/index.html"
	// http.ServeFile(w, r, home)
}
