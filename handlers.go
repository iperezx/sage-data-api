package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type nodeSchema struct {
	NodeID        string `json:"nodeID,omitempty"`
	MetadataName  string `json:"metadataName,omitempty"`
	MetadataValue string `json:"metadataValue,omitempty"`
}

type nodeSage struct {
	ID     string `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Status string `json:"status,omitempty"`
	Lat    string `json:"lat,omitempty"`
	Lon    string `json:"lon,omitempty"`
}

func getSageNodes(w http.ResponseWriter, r *http.Request) {
	data := []*nodeSage{}
	data = getAllSageNodes()
	log.Println("GET All Nodes")
	respondJSON(w, http.StatusOK, data)
	return
}

func postSageNodes(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	dataOut := new(nodeSchema)
	dataOut.NodeID = query.Get("nodeid")
	dataOut.MetadataName = query.Get("metadata_name")
	dataOut.MetadataValue = query.Get("metadata_value")

	insertNode(dataOut)

	log.Println("POST(INSERT): ")
	log.Printf("nodeID : %s", dataOut.NodeID)
	log.Printf("metadataName : %s", dataOut.MetadataName)
	log.Printf("metadataValue : %s", dataOut.MetadataValue)
	log.Println()

	respondJSON(w, http.StatusOK, dataOut)
	return
}

func getSageNode(nodeID string) []*nodeSchema {
	db, err := sql.Open("mysql", mysqlDSN)
	if err != nil {
		err = fmt.Errorf("Unable to connect to database: %v", err)
		return nil
	}
	defer db.Close()
	queryStr := "SELECT * FROM Nodes where nodeid=? ;"
	stmt, err := db.Prepare(queryStr)

	if err != nil {
		err = fmt.Errorf("DB Prepare Error: %v", err)
		return nil
	}

	nodeData, err := stmt.Query(nodeID)

	if err != nil {
		err = fmt.Errorf("Query Error: %v", err)
		return nil
	}
	dataOut := []*nodeSchema{}
	for nodeData.Next() {
		row := new(nodeSchema)
		err = nodeData.Scan(&row.NodeID, &row.MetadataName, &row.MetadataValue)
		if err != nil {
			err = fmt.Errorf("Error with parsing row: %v", err)
			return nil
		}
		dataOut = append(dataOut, row)
	}
	return dataOut
}

func getAllSageNodes() []*nodeSage {
	db, err := sql.Open("mysql", mysqlDSN)
	if err != nil {
		err = fmt.Errorf("Unable to connect to database: %v", err)
		return nil
	}
	defer db.Close()
	queryStr := "SELECT DISTINCT(nodeid) FROM Nodes ;"
	stmt, err := db.Prepare(queryStr)

	if err != nil {
		err = fmt.Errorf("DB Prepare Error: %v", err)
		return nil
	}

	nodeIDs, err := stmt.Query()

	if err != nil {
		err = fmt.Errorf("Query Error: %v", err)
		return nil
	}

	dataOut := []*nodeSage{}
	dataOut = getNodeIDRecords(nodeIDs)
	return dataOut
}

func getNodeIDRecords(nodeIDs *sql.Rows) []*nodeSage {
	dataOut := []*nodeSage{}
	for nodeIDs.Next() {
		var nodeID string
		err := nodeIDs.Scan(&nodeID)
		if err != nil {
			err = fmt.Errorf("Error with parsing row: %v", err)
			return nil
		}

		var nodeData []*nodeSchema
		nodeData = getSageNode(nodeID)
		nodeOut := new(nodeSage)
		for _, node := range nodeData {
			nodeOut.ID = node.NodeID
			if node.MetadataName == "name" {
				nodeOut.Name = node.MetadataValue
			}
			if node.MetadataName == "status" {
				nodeOut.Status = node.MetadataValue
			}
			if node.MetadataName == "lat" {
				nodeOut.Lat = node.MetadataValue
			}
			if node.MetadataName == "lon" {
				nodeOut.Lon = node.MetadataValue
			}
		}
		dataOut = append(dataOut, nodeOut)
	}
	return dataOut
}

func insertNode(node *nodeSchema) {
	db, err := sql.Open("mysql", mysqlDSN)
	if err != nil {
		err = fmt.Errorf("Unable to connect to database: %v", err)
		return
	}
	defer db.Close()
	insertQueryStr := "INSERT INTO Nodes (nodeid, metadata_name, metadata_value) VALUES ( ?, ?, ?)  ;"
	insForm, err := db.Prepare(insertQueryStr)
	if err != nil {
		err = fmt.Errorf("Node insertion prepare in mysql failed: %s", err.Error())
		return
	}
	insForm.Exec(node.NodeID, node.MetadataName, node.MetadataValue)
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
