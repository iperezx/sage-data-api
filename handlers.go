package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type nodeSage struct {
	ID             string `json:"id,omitempty"`
	Name           string `json:"name,omitempty"`
	Status         string `json:"status,omitempty"`
	ProvisionDate  string `json:"provisionDate,omitempty"`
	OSVersion      string `json:"OSVersion,omitempty"`
	ServiceTag     string `json:"serviceTag,omitempty"`
	SpecialDevices string `json:"SpecialDevices,omitempty"`
	BiosVersion    string `json:"BiosVersion,omitempty"`
	Lat            string `json:"lat,omitempty"`
	Lon            string `json:"lon,omitempty"`
}

func getSageNodes(w http.ResponseWriter, r *http.Request) {
	data := getNodeDataFromCSV(csvFile)
	log.Println("GET All Nodes")
	respondJSON(w, http.StatusOK, data)
	return
}

func getNodeDataFromCSV(csvFile string) []*nodeSage {
	csv_file, err := os.Open(csvFile)
	if err != nil {
		err = fmt.Errorf("Error with opening csv: %v", err)
		return nil
	}
	defer csv_file.Close()

	r := csv.NewReader(csv_file)
	records, err := r.ReadAll()
	if err != nil {
		err = fmt.Errorf("Error with reading all data: %v", err)
		return nil
	}
	header := make([]string, 0)
	for _, headerContent := range records[0] {
		header = append(header, headerContent)
	}
	indexMap := map[string]int{}
	for i, name := range header {
		indexMap[name] = i
	}
	records = records[1:]
	nodes := []*nodeSage{}
	for _, rec := range records {
		node := new(nodeSage)
		node.ID = rec[indexMap["ID"]]
		node.Name = rec[indexMap["Name"]]
		node.Status = rec[indexMap["Status"]]
		node.ProvisionDate = rec[indexMap["ProvisionDate"]]
		node.OSVersion = rec[indexMap["OSVersion"]]
		node.ServiceTag = rec[indexMap["ServiceTag"]]
		node.SpecialDevices = rec[indexMap["SpecialDevices"]]
		node.BiosVersion = rec[indexMap["BiosVersion"]]
		node.Lat = rec[indexMap["Lat"]]
		node.Lon = rec[indexMap["Lon"]]
		nodes = append(nodes, node)
	}
	return nodes
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
