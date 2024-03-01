package main

import (
	"fmt"
	// "encoding/json"
	// "io"
	// "net/http"
	// "log"
)

type Human struct{
	Name string `json:"human_name"`
	Age int
}

func print[T string | int | float64 ](value T){
	fmt.Printf("There is %#v", value)
}


func main(){
	// fmt.Printf("Test")

	// client := &http.Client{}
	// resp, err := client.Get("http://theveloper.pl")

	// if err != nil{
	// 	fmt.Errorf("Error occured %d", err)
	// }

	// defer resp.Body.Close()
	// body, err := io.ReadAll(resp.Body)
	// if err != nil{
	// 	fmt.Errorf("Error occured %d", err)
	// }

	// sb := string(body)
	// log.Printf(sb)


	// p1 := Human{Name: "Mateusz", Age:12}
	// p1bytes, err := json.Marshal(p1)
	// if err != nil{
	// 	fmt.Errorf("Error occured %d", err)
	// }

	// fmt.Printf(string(p1bytes))
	// fmt.Print(p1bytes)

	print("XD")
	print(2134)
	print(2134.23)


}