package main

import (
	"fmt"
	"log"
	"net/http"

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
	csvFile       string
	jsonFile      string
)

func init() {
	csvFile = "manifest.csv"
	jsonFile = "dataDict.json"
}

func createRouter() {

	mainRouter = mux.NewRouter()
	r := mainRouter
	log.Println("Sage Node API")
	api := r.PathPrefix("/api/v1").Subrouter()

	api.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to SAGE Node API")
	})

	// GET /nodes
	api.Handle("/nodes-data", negroni.New(
		negroni.HandlerFunc(authMW),
		negroni.Wrap(http.HandlerFunc(getSageNodes)),
	)).Methods(http.MethodGet)

	api.Handle("/nodes-metadata", negroni.New(
		negroni.HandlerFunc(authMW),
		negroni.Wrap(http.HandlerFunc(getSageNodesMetadata)),
	)).Methods(http.MethodGet)

	api.Handle("/nodes-all", negroni.New(
		negroni.HandlerFunc(authMW),
		negroni.Wrap(http.HandlerFunc(getSageNodesAndDataDict)),
	)).Methods(http.MethodGet)

	log.Fatalln(http.ListenAndServe(":8080", api))

}

func main() {
	createRouter()
}
