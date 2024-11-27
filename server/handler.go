package server

import (
	"animal-hash-digest/generator"
	"encoding/json"
	"net/http"
	"strconv"
)

// HTTP handler for the /hostname endpoint
func hostnameHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the "count" query parameter (default: 1)
	count := 1
	queryCount := r.URL.Query().Get("count")
	if queryCount != "" {
		if parsedCount, err := strconv.Atoi(queryCount); err == nil && parsedCount > 0 {
			count = parsedCount
		}
	}

	// Generate hostnames
	hostnames := make([]string, count)
	for i := 0; i < count; i++ {
		hostnames[i] = generator.GenerateHostname()
	}

	// Return the hostnames as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"hostnames": hostnames,
	})
}
