// Golang program to illustrate the
// concept of html/templates
package main

import (
	"fmt"
	"html/template"
	"net/http"
	// "os"
	"reflect"
)


type Person struct {
    UserName string
	Emails []string
}

func hello(w http.ResponseWriter, req *http.Request){
    t, err := template.New("index.html").ParseFiles("index.html")
	fmt.Print(err)
    p := Person{UserName: "Astaxie", Emails:  []string{"astaxie@beego.me", "astaxie@gmail.com"}}

	t.Execute(w, p)
    e := reflect.ValueOf(&p).Elem()
	
    for i := 0; i < e.NumField(); i++ {
		varName := e.Type().Field(i).Name
		varType := e.Type().Field(i).Type
		varValue := e.Field(i).Interface()
		fmt.Fprintf(w, "%v %v %v\n", varName,varType,varValue)
    }

}

func main() {
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8090", nil)
}