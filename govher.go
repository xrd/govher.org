package main

import (
	"fmt"
	"net/http"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	path := ""
	path = r.URL.Path[:]
	// Replace github with github.com
	path = strings.Replace( path, "github", "github.com", 1 )
	fmt.Println( "Path: " + path )
	http.Redirect(w, r, "http://" + path, http.StatusFound)
        return
	// fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":3000", nil)
}

