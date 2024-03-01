package main

import (
	"strings"
	"io"
	"os"
	"testing"
	"fmt"
)


func Test(t *testing.T){

	var msg = []string{"is seated at the table.", "takes the right fork.", "takes the left fork.", "has both forks and is eating.", "is thinking.", "put down the forks.", "is satisfied.", "left the table."}
	stdOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	_ = w.Close()
	result, _ := io.ReadAll(r)
	output := string(result)
	os.Stdout = stdOut



	for i:=0; i<len(philosophers); i++{
		for j:=0; j<len(msg); j++{
			new := fmt.Sprintf("%v %v",philosophers[i].name, msg[j])
			if !strings.Contains(output, new){
				t.Errorf("Expected: %s, Recived: %s",new, output)
			}
		}
	}

}


