package main

import (
	"database/sql"
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/alexedwards/scs/redisstore"
	"github.com/alexedwards/scs/v2"
	"github.com/gomodule/redigo/redis"
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/lib/pq"
)

const webPort = "80"
const POSTGRES_USER = "postgres"
const POSTGRES_PASSWORD = "password"
const POSTGRES_DB = "concurrency"
const REDIS = "127.0.0.1:6379"
func main(){
	db := initDB()
	db.Ping()
	session := initSession()

	wg := sync.WaitGroup{}

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)


	app := Config{
		Session: session,
		DB: db,
		InfoLog: infoLog,
		ErrorLog: errorLog,
		Wait: &wg,
		Models: data.New(db),
		ErrorChan: make(chan error),
		ErrorChanDone: make(chan bool),
	}

	app.Mailer = app.createMail()
	go app.listenForMail()

	go app.listenForErrors()

	app.serve()
}

func (app *Config) serve(){
	srv := &http.Server{
		Addr: fmt.Sprintf(":%s",webPort),
		Handler: app.routes(),
	}
	app.InfoLog.Println("Starting web server...")
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}

func initDB() *sql.DB{
	conn := connectToDB()
	if conn == nil {
		log.Panic("Cant connect to database")
	}
	return conn
}

func connectToDB() *sql.DB{
	counts :=0

	for {

		if counts > 5 {
			return nil
		}

		connection, err := openDB()
		if err != nil {
			log.Println("Postgres not yet ready...")
		} else {
			log.Println("Connected to database !")
			return connection
		}
		counts += 1
		time.Sleep(time.Second * 2)
	}

}

func openDB() (*sql.DB, error){
	connStr := fmt.Sprintf("postgresql://%s:%s@0.0.0.0:5432/%s?sslmode=disable",POSTGRES_USER,POSTGRES_PASSWORD, POSTGRES_DB)
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Println(err)
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return db, nil
}

func initSession()* scs.SessionManager {
	gob.Register(data.User{})
	session := scs.New()
	session.Store = redisstore.New(initRedis())
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = true
	return session
}

func initRedis() *redis.Pool {
	redisPool := &redis.Pool{
		MaxIdle: 10,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp",REDIS)
		},
	}
	return redisPool
}

func (app *Config) createMail() Mail {

	errorChan := make(chan error)
	mailerChan := make(chan Message, 100)
	mailerDoneChan := make(chan bool)

	m := Mail{
		Domain: "localhost",
		Host: "localhost",
		Port: 1025,
		Encryption: "none",
		FromAddress: "info@mycompany.com",
		FromName: "info",
		ErrorChan: errorChan,
		MailerChan: mailerChan,
		DoneChan: mailerDoneChan,
		Wait: app.Wait,

	}
	return m
}

func (app *Config) listenForErrors() {
	for {
		select {
		case err := <- app.ErrorChan:
			app.ErrorLog.Println(err)
		case <- app.ErrorChanDone:
			return
		}
	}
}