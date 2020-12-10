package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	geo "github.com/paulmach/go.geo"
)

type sageNode struct {
	NodeID        string    `json:"nodeID,omitempty"`
	MetadataName  string    `json:"metadataName,omitempty"`
	MetadataValue string    `json:"metadataValue,omitempty"`
	Geom          geo.Point `json:"geom,omitempty"`
}

func getSageNodes(w http.ResponseWriter, r *http.Request) {
	dataOut := []*sageNode{}
	dataOut = getAllSageNodes()
	log.Println("GET All Nodes")
	respondJSON(w, http.StatusOK, dataOut)
	return
}

func postSageNodes(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	dataOut := new(sageNode)
	dataOut.NodeID = query.Get("nodeid")
	dataOut.MetadataName = query.Get("metadata_name")
	dataOut.MetadataValue = query.Get("metadata_value")
	lonStr := query.Get("lon")
	latStr := query.Get("lat")
	lon, err := strconv.ParseFloat(lonStr, 64)
	if err != nil {
		err = fmt.Errorf("Error converting string to float: %v", err)
		return
	}
	lat, err := strconv.ParseFloat(latStr, 64)
	if err != nil {
		err = fmt.Errorf("Error converting string to float: %v", err)
		return
	}
	dataOut.Geom = *geo.NewPoint(lat, lon)
	insertNode(dataOut)

	log.Println("POST(INSERT): ")
	log.Printf("nodeID : %s", dataOut.NodeID)
	log.Printf("metadataName : %s", dataOut.MetadataName)
	log.Printf("metadataValue : %s", dataOut.MetadataValue)
	log.Printf("geom : %s", dataOut.Geom.ToWKT())
	log.Printf("\n")

	respondJSON(w, http.StatusOK, dataOut)
	return
}

func getAllSageNodes() []*sageNode {
	db, err := sql.Open("mysql", mysqlDSN)
	if err != nil {
		err = fmt.Errorf("Unable to connect to database: %v", err)
		return nil
	}
	defer db.Close()
	queryStr := "SELECT * FROM Nodes ;"
	stmt, err := db.Prepare(queryStr)

	if err != nil {
		err = fmt.Errorf("DB Prepare Error: %v", err)
		return nil
	}

	data, err := stmt.Query()

	if err != nil {
		err = fmt.Errorf("Query Error: %v", err)
		return nil
	}

	dataOut := []*sageNode{}
	for data.Next() {
		row := new(sageNode)
		err = data.Scan(&row.NodeID, &row.MetadataName, &row.MetadataValue, &row.Geom)
		if err != nil {
			err = fmt.Errorf("Error with parsing row: %v", err)
			return nil
		}
		dataOut = append(dataOut, row)
	}
	return dataOut
}

func insertNode(node *sageNode) {
	db, err := sql.Open("mysql", mysqlDSN)
	if err != nil {
		err = fmt.Errorf("Unable to connect to database: %v", err)
		return
	}
	defer db.Close()
	insertQueryStr := "INSERT INTO Nodes (nodeid, metadata_name, metadata_value, geom) VALUES ( ?, ?, ?, ST_GeomFromText( ? , 4326))  ;"
	insForm, err := db.Prepare(insertQueryStr)
	if err != nil {
		err = fmt.Errorf("Node insertion prepare in mysql failed: %s", err.Error())
		return
	}
	insForm.Exec(node.NodeID, node.MetadataName, node.MetadataValue, node.Geom.ToWKT())
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
