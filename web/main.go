package main

import (
    "fmt"
    "net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {

    fmt.Fprintf(w, "hello\n")
	fmt.Printf("Hello")
}

func headers(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, "hello\n")
	fmt.Printf("Hello")
    for name, headers := range req.Header {
        for _, h := range headers {
            fmt.Fprintf(w, "%v: %v\n", name, h)
        }
    }
    main2(w)
}

func main() {

    http.HandleFunc("/hello", hello)
    http.HandleFunc("/headers", headers)
    http.Handle("/", http.FileServer(http.Dir("./static")))
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./public"))))

    http.ListenAndServe(":8090", nil)
}