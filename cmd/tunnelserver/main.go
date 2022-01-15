package main

import (
	"fmt"
	"math/rand"
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

	// c := 10
	// b := make([]byte, c)
	// // _, err := rand.Read(b)
	// // if err != nil {
	// // 	fmt.Println("error:", err)
	// // 	return
	// // }

	// if _, err := io.ReadFull(rand.Reader, b); err != nil {
	// 		fmt.Println("error:", err)
	// 		http.Error(w, "Failed stream.Send", http.StatusInternalServerError)
	// 	return
	// }
	// fmt.Println("random str", hex.EncodeToString(b))

	rand.Seed(time.Now().UnixNano())
	rnd := rand.Intn(100000)
	fmt.Println("random int", rnd)

	<-time.After(10 * time.Second)

	// apiURL, _ := url.Parse("https://gitlab.com")
	// fmt.Println("url.Parse", apiURL.Scheme)

	fmt.Fprint(w, r.URL.String())
}

func HelloServer(w http.ResponseWriter, r *http.Request) {

	// r.URL.String() or r.URL.RequestURI() show only r.URL.Path!
	// r.URL.Scheme is empty! Possible  r.TLS == nil
	// FullUrl := "https://" + r.Host + r.URL.Path

	// for key, values := range r.Header {
	// 	for _, val := range values {
	// 		w.Header().Add(key, val)
	// 	}
	// }
	// w.WriteHeader(w.Status) //if you have w.WriteHeader(200) before w.Header the headers will not be set!!!

	fmt.Println("Server:URL", r.URL.String())
	fmt.Println("Server:RequestURI", r.URL.RequestURI())
	fmt.Println("Server:URL.Path", r.URL.Path)
	fmt.Println("Server:URL.Scheme", r.URL.Scheme)
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])

	// c := 10
	// b := make([]byte, c)
	// // _, err := rand.Read(b)
	// // if err != nil {
	// // 	fmt.Println("error:", err)
	// // 	return
	// // }

	// if _, err := io.ReadFull(rand.Reader, b); err != nil {
	// 		fmt.Println("error:", err)
	// 		http.Error(w, "Failed stream.Send", http.StatusInternalServerError)
	// 	return
	// }
	// fmt.Println("random str", hex.EncodeToString(b))

	rand.Seed(time.Now().UnixNano())
	rnd := rand.Intn(100)
	fmt.Println("random int", rnd)

	<-time.After(10 * time.Second)

	// apiURL, _ := url.Parse("https://gitlab.com")
	// fmt.Println("url.Parse", apiURL.Scheme)

	fmt.Fprint(w, r.URL.String())
}

func main() {
	h1 := httpHandler{message: "Index"}

	http.Handle("/", h1)

	// http.HandleFunc("/", HelloServer)

	fmt.Println("ListenAndServe start!")
	http.ListenAndServe(":8090", nil)

}
