package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	geo "github.com/paulmach/go.geo"
)

type sqlSchema struct {
	NodeID        string    `json:"nodeID,omitempty"`
	MetadataName  string    `json:"metadataName,omitempty"`
	MetadataValue string    `json:"metadataValue,omitempty"`
	Geom          geo.Point `json:"geom,omitempty"`
}

func getSageNodes(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", mysqlDSN)
	if err != nil {
		err = fmt.Errorf("Unable to connect to database: %v", err)
		return
	}
	defer db.Close()
	queryStr := "SELECT * FROM Nodes ;"
	stmt, err := db.Prepare(queryStr)

	if err != nil {
		err = fmt.Errorf("DB Prepare Error: %v", err)
		return
	}

	data, err := stmt.Query()

	if err != nil {
		err = fmt.Errorf("Query Error: %v", err)
		return
	}

	dataOut := []*sqlSchema{}
	for data.Next() {
		row := new(sqlSchema)
		err = data.Scan(&row.NodeID, &row.MetadataName, &row.MetadataValue, &row.Geom)
		if err != nil {
			err = fmt.Errorf("Error with parsing row: %v", err)
			return
		}
		dataOut = append(dataOut, row)
	}

	log.Printf("GetSageNodes, queryStr: %s", queryStr)
	respondJSON(w, http.StatusOK, dataOut)
	return
}

func authMW(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	// Add sage authentication later (if needed)
	next(w, r)
}

func respondJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	s, err := json.MarshalIndent(data, "", "  ")
	if err == nil {
		w.Write(s)
	}
}
