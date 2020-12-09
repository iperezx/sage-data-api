package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

var (
	mysqlHost     string
	mysqlDatabase string
	mysqlUsername string
	mysqlPassword string
	mysqlDSN      string // Data Source Name
	mainRouter    *mux.Router
)

func init() {

	mysqlHost = os.Getenv("MYSQL_HOST")
	mysqlDatabase = os.Getenv("MYSQL_DATABASE")
	mysqlUsername = os.Getenv("MYSQL_USER")
	mysqlPassword = os.Getenv("MYSQL_PASSWORD")

	mysqlDSN = fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", mysqlUsername, mysqlPassword, mysqlHost, mysqlDatabase)

	log.Printf("mysqlHost: %s", mysqlHost)
	log.Printf("mysqlDatabase: %s", mysqlDatabase)
	log.Printf("mysqlUsername: %s", mysqlUsername)
	log.Printf("mysqlDSN: %s", mysqlDSN)
	count := 0
	for {
		count++
		db, err := sql.Open("mysql", mysqlDSN)
		if err != nil {
			if count > 1000 {
				log.Fatalf("(sql.Open) Unable to connect to database: %v", err)
				return
			}
			log.Printf("(sql.Open) Unable to connect to database: %v, retrying...", err)
			time.Sleep(time.Second * 3)
			continue
		}
		//err = db.Ping()
		for {
			_, err = db.Exec("DO 1")
			if err != nil {
				if count > 1000 {
					log.Fatalf("(db.Ping) Unable to connect to database: %v", err)
					return
				}
				log.Printf("(db.Ping) Unable to connect to database: %v, retrying...", err)
				time.Sleep(time.Second * 3)
				continue
			}
			break
		}
		break
	}
}

func createRouter() {

	mainRouter = mux.NewRouter()
	r := mainRouter
	log.Println("Sage Node API")
	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to SAGE Node API")
	})

	// - list buckets
	// GET /objects/
	api.Handle("/nodes", negroni.New(
		negroni.HandlerFunc(authMW),
		negroni.Wrap(http.HandlerFunc(getSageNodes)),
	)).Methods(http.MethodGet)

	log.Fatalln(http.ListenAndServe(":8080", api))

}

func main() {
	createRouter()
}
