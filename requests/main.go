package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)



type logWriter struct{

}

func main() {
	res, err := http.Get("http://localhost:8090/headers")
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", body)
}

func (logWriter) Write(bs []byte) (int,error){
	return 1,nil
}