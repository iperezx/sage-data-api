package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type nodeSage struct {
	Name           string `json:"name,omitempty"`
	ID             string `json:"id,omitempty"`
	Status         string `json:"status,omitempty"`
	ProvisionDate  string `json:"provisionDate,omitempty"`
	OSVersion      string `json:"OSVersion,omitempty"`
	ServiceTag     string `json:"serviceTag,omitempty"`
	SpecialDevices string `json:"SpecialDevices,omitempty"`
	BiosVersion    string `json:"BiosVersion,omitempty"`
	Lat            string `json:"lat,omitempty"`
	Lon            string `json:"lon,omitempty"`
}

type metadata struct {
	ID          string `json:"id,omitempty"`
	Type        string `json:"type,omitempty"`
	Label       string `json:"label,omitempty"`
	Description string `json:"description,omitempty"`
}

type data struct {
	Data     []*nodeSage `json:"data,omitempty"`
	Metadata []*metadata `json:"metadata,omitempty"`
}

type pluginMetadata struct {
	Node     string `json:"node,omitempty"`
	Host     string `json:"host,omitempty"`
	Plugin   string `json:"plugin,omitempty"`
	Camera   string `json:"camera,omitempty"`
	SensorID string `json:"sensorID,omitempty"`
}

type pluginSage struct {
	Timestamp string         `json:"timestamp,omitempty"`
	Name      string         `json:"name,omitempty"`
	Value     float64        `json:"value,omitempty"`
	Metadata  pluginMetadata `json:"metadata,omitempty"`
}

type pluginData struct {
	Data     []*pluginSage `json:"data"`
	Metadata []*metadata   `json:"metadata"`
}

type sensorHardware struct {
	ID           string `json:"id,omitempty"`
	Product_name string `json:"product_name,omitempty"`
	Manufacture  string `json:"manufacture,omitempty"`
	Sensor_types string `json:"sensor_types,omitempty"`
	Link         string `json:"link,omitempty"`
}

type sensorHardwareData struct {
	Data     []*sensorHardware `json:"data"`
	Metadata []*metadata       `json:"metadata"`
}

func getSageNodes(w http.ResponseWriter, r *http.Request) {
	nodeData := getNodeDataFromCSV(csvFile)
	var allData data
	allData.Data = nodeData
	log.Println("GET: All nodes data")
	respondJSON(w, http.StatusOK, allData)
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
		node.Name = rec[indexMap["name"]]
		node.ID = rec[indexMap["id"]]
		node.Status = rec[indexMap["status"]]
		node.ProvisionDate = rec[indexMap["provisionDate"]]
		node.OSVersion = rec[indexMap["OSVersion"]]
		node.ServiceTag = rec[indexMap["ServiceTag"]]
		node.SpecialDevices = rec[indexMap["SpecialDevices"]]
		node.BiosVersion = rec[indexMap["BIOSVersion"]]
		node.Lat = rec[indexMap["Lat"]]
		node.Lon = rec[indexMap["Lon"]]
		nodes = append(nodes, node)
	}
	return nodes
}

func getSageNodesMetadata(w http.ResponseWriter, r *http.Request) {
	nodeMetadata := getNodeMetadataFromJson(jsonFile)
	var allData data
	allData.Metadata = nodeMetadata
	log.Println("GET: All nodes metadata")
	respondJSON(w, http.StatusOK, allData)
	return
}

func getNodeMetadataFromJson(jsonFile string) []*metadata {
	jsonData, err := os.Open(jsonFile)
	if err != nil {
		fmt.Println(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonData)
	data := []*metadata{}
	json.Unmarshal(byteValue, &data)
	return data
}

func getSageNodesAndDataDict(w http.ResponseWriter, r *http.Request) {
	nodeData := getNodeDataFromCSV(csvFile)
	metadata := getNodeMetadataFromJson(jsonFile)
	var allData data
	allData.Data = nodeData
	allData.Metadata = metadata
	log.Println("GET: All nodes data and metadata")
	respondJSON(w, http.StatusOK, allData)
	return
}

//Plugin Data section

func getPluginDataFromJson(jsonFile string) []*pluginSage {
	jsonData, err := os.Open(jsonFile)
	if err != nil {
		fmt.Println(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonData)
	data := []*pluginSage{}
	json.Unmarshal(byteValue, &data)
	return data
}

func getSagePluginData(w http.ResponseWriter, r *http.Request) {
	data := getPluginDataFromJson(pluginFile)
	var sagePluginData pluginData
	sagePluginData.Data = data
	sagePluginData.Metadata = []*metadata{}
	log.Println("GET: All plugin data")
	respondJSON(w, http.StatusOK, sagePluginData)
	return
}

//Sensor Hardware Data Section
func getSensorHardwareFromJson(jsonFile string) []*sensorHardware {
	jsonData, err := os.Open(jsonFile)
	if err != nil {
		fmt.Println(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonData)
	data := []*sensorHardware{}
	json.Unmarshal(byteValue, &data)
	return data
}

func getSensorHardwareData(w http.ResponseWriter, r *http.Request) {
	data := getSensorHardwareFromJson(sensorHardwareFile)
	var sensorHardware sensorHardwareData
	sensorHardware.Data = data
	sensorHardware.Metadata = []*metadata{}
	log.Println("GET: All sensor hardware data")
	respondJSON(w, http.StatusOK, sensorHardware)
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
