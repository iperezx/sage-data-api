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
	wildNodesFile      string
	pluginFile         string
	sensorHardwareFile string
)

func init() {
	csvFile = "init_data/manifest.csv"
	jsonFile = "init_data/dataDict.json"
	wildNodesFile = "init_data/wild_nodes.json"
	pluginFile = "init_data/pluginData.json"
	sensorHardwareFile = "init_data/sensor_hardware.json"
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
	// /nodes blades
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

	// /nodes-wild
	api.Handle("/nodes-wild-data", negroni.New(
		negroni.HandlerFunc(authMW),
		negroni.Wrap(http.HandlerFunc(getWildNodeData)),
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
