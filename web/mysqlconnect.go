package main

import (
    "database/sql"
    "fmt"
    "log"
    // "os"
    "net/http"
    "github.com/go-sql-driver/mysql"
)


var db *sql.DB

type City struct {
    Id         int
    Name       string
    Population int
}


func main2(w http.ResponseWriter) {
    // Capture connection properties.
    cfg := mysql.Config{
        User:   "root",
        Passwd: "Nokia123123!",
        Net:    "tcp",
        Addr:   "127.0.0.1:3306",
        DBName: "testy",
		AllowNativePasswords: true,
    }
    // Get a database handle.
    var err error
    db, err = sql.Open("mysql", cfg.FormatDSN())
    if err != nil {
        log.Fatal(err)
    }

    pingErr := db.Ping()
    if pingErr != nil {
        log.Fatal(pingErr)
    }
    fmt.Println("Connected!")


    res, err := db.Query("SELECT * FROM cities")

    defer res.Close()

    if err != nil {
        log.Fatal(err)
    }

    cols, _ := res.Columns()
    fmt.Print(cols)

    for res.Next() {

        var city City
        err := res.Scan(&city.Id, &city.Name, &city.Population)

        if err != nil {
            log.Fatal(err)
        }

        // fmt.Printf("%v\n", city)
		fmt.Fprintf(w, "%v", city)
    }


}