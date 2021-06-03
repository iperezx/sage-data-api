package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

var (
	mainRouter         *mux.Router
	csvFile            string
	jsonFile           string
	pluginFile         string
	sensorHardwareFile string
)

func init() {
	csvFile = "manifest.csv"
	jsonFile = "dataDict.json"
	pluginFile = "pluginData.json"
	sensorHardwareFile = "sensor_hardware.json"
}

func createRouter() {

	mainRouter = mux.NewRouter()
	r := mainRouter
	log.Println("Sage Data API")
	api := r.PathPrefix("/api/v1").Subrouter()

	api.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to SAGE Data API")
	})

	// GET
	// /nodes
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

	// /plugin
	api.Handle("/plugin-data", negroni.New(
		negroni.HandlerFunc(authMW),
		negroni.Wrap(http.HandlerFunc(getSagePluginData)),
	)).Methods(http.MethodGet)

	// /sensor-hardware
	api.Handle("/sensor-hardware-data", negroni.New(
		negroni.HandlerFunc(authMW),
		negroni.Wrap(http.HandlerFunc(getSensorHardwareData)),
	)).Methods(http.MethodGet)

	log.Fatalln(http.ListenAndServe(":8080", api))

}

func main() {
	createRouter()
}
