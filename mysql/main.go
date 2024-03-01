package main

import (
	"database/sql"
	"fmt"
	"log"
	// "os"
    "reflect"
	"github.com/go-sql-driver/mysql"
)


var db *sql.DB

type City struct {
    Id         int
    Name       string
    Population int
}

func main(){
    main2("Warsaw")
}

func connectDb(cfg mysql.Config) (db *sql.DB, err error){
    db, err = sql.Open("mysql", cfg.FormatDSN())
    if err != nil {
        log.Fatal(err)
    }

    return db,err
}

func selectWithCondition(db *sql.DB, from string, conditionName string, condition string) (res *sql.Rows, err error){
    res, err = db.Query("SELECT * FROM "+from+" WHERE "+conditionName+" = ?", condition)

    return res,err
}

func main2(condition string) {
    cfg := mysql.Config{
        User:   "root",
        Passwd: "Nokia123123!",
        Net:    "tcp",
        Addr:   "127.0.0.1:3306",
        DBName: "testy",
		AllowNativePasswords: true,
    }
    // Get a database handle.
    db, err := connectDb(cfg)


    pingErr := db.Ping()
    if pingErr != nil {
        log.Fatal(pingErr)
    }
    fmt.Println("Connected!")


    res, err := selectWithCondition(db, "cities", "name", "Warsaw")

    defer res.Close()

    if err != nil {
        log.Fatal(err)
    }


    // for res.Next() {

    //     var city City
    //     err := res.Scan(&city.Id, &city.Name, &city.Population)

    //     if err != nil {
    //         log.Fatal(err)
    //     }

    //     fmt.Printf("%v\n", city)
    // }

    for res.Next() {
        user := City{}

        s := reflect.ValueOf(&user).Elem()
        numCols := s.NumField()
        columns := make([]interface{}, numCols)
        for i := 0; i < numCols; i++ {
            field := s.Field(i)
            columns[i] = field.Addr().Interface()
        }

        err := res.Scan(columns...)
        if err != nil {
            log.Fatal(err)
        }
        log.Println(user)
    }




}