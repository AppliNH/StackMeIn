package apihandler

import (
	"encoding/json"
	"fmt"
	"net/http"

	t "primitivo.fr/applinh/go-docker-compose/dc_file_mng"
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

	if uuid, erro := t.ParseComposeData(data["version"].(string), data["services"].(map[string]interface{}), data["networks"].(map[string]interface{})); erro != nil {
		fmt.Println(erro)
		http.Error(w, erro.Error(), http.StatusBadRequest)
		return
	} else {
		response := map[string]string{"statusCode": "200", "success": "true", "id": uuid}
		json.NewEncoder(w).Encode(response)
	}
}
