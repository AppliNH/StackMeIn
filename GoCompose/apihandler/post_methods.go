package apihandler

import (
	"encoding/json"
	"fmt"
	"net/http"

	db "primitivo.fr/applinh/go-docker-compose/dbcontroller"
)

func POST_dockercompose(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data := make(map[string]interface{})
	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	write_data := make(map[string]interface{})
	write_data["dockercompose"] = data
	write_data["containers"] = []string{}

	dbRes, err := db.WriteToDB("stacks", write_data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else {
		response := map[string]string{"statusCode": "200", "id": dbRes["id"].(string)}
		json.NewEncoder(w).Encode(response)
	}
}
