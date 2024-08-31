package api

import (
	"encoding/json"
	"museum/go/server/data"
	"net/http"
	"strconv"
)

func Get(w http.ResponseWriter, r *http.Request) { 
	// sending json directly 
	w.Header().Set("Content-Type", "application/json")
	// api/exhibitions?id=777
	id := r.URL.Query()["id"]

	if id != nil { 
		finalId, err := strconv.Atoi(id[0])

		if err == nil && finalId < len(data.GetAll()) { 
			json.NewEncoder(w).Encode(data.GetAll()[finalId])
		} else { 
			http.Error(w, "Invalid id", http.StatusBadRequest)
		}
	} else { 
		json.NewEncoder(w).Encode(data.GetAll())
	}
	
}