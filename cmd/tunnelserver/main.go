package main

import (
	"fmt"
	"net/http"
	"time"
)

type httpHandler struct {
	message string
}

func (h httpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Server:URL", r.URL.String())
	fmt.Println("Server:RequestURI", r.URL.RequestURI())
	fmt.Println("Server:URL.Path", r.URL.Path)
	fmt.Println("Server:URL.Scheme", r.URL.Scheme)
	fmt.Println("Server:Host", r.Host)
	fmt.Println("Server:URL.Host", r.URL.Host)
	fmt.Println("Server:Header", r)

	<-time.After(10 * time.Second)

	// apiURL, _ := url.Parse("https://gitlab.com")
	// fmt.Println("url.Parse", apiURL.Scheme)

	fmt.Fprint(w, r.URL.String())
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Server:URL", r.URL.String())
	fmt.Println("Server:RequestURI", r.URL.RequestURI())
	fmt.Println("Server:URL.Path", r.URL.Path)
	fmt.Println("Server:URL.Scheme", r.URL.Scheme)
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

func main() {
	h1 := httpHandler{message: "Index"}

	http.Handle("/", h1)

	fmt.Println("ListenAndServe start!")

	// http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8090", nil)

}
