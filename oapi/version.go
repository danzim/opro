package oapi

import (
	"encoding/json"
	"net/http"
)

type structVersion struct {
	Version string `json:"version"`
}

// GetVersion - Function to expose Version of API
func GetVersion(w http.ResponseWriter, r *http.Request) {
	version := &structVersion{
		Version: "0.1",
	}
	json.NewEncoder(w).Encode(version)
}
