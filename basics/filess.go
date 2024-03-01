package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main3() {
    f, err := os.Open("app.txt")

    if err != nil {
        log.Fatal(err)
    }

    defer f.Close()

    scanner := bufio.NewScanner(f)

    for scanner.Scan() {
		line := scanner.Text()
		items := strings.Split(line, " ")
		fmt.Printf("%#v",items)
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}