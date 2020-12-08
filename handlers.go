package main

import (
	"fmt"
	"net/http"
)

// getSageNodes
func getSageNodes(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Nodes:\n")
	w.WriteHeader(http.StatusOK)
	return
}

func authMW(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	// Add sage authentication later (if needed)
	next(w, r)
}
