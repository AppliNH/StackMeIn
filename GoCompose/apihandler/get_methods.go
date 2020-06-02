package apihandler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	t "primitivo.fr/applinh/go-docker-compose/dc_file_mng"
	d "primitivo.fr/applinh/go-docker-compose/dockercontrol"
)

func GET_Stack(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if containerID, err1 := d.CreateNewContainer(id); err1 == nil {
		containers, _ := d.ListContainers()
		found := false
		for _, v := range containers {
			if v == containerID {
				found = true
				json.NewEncoder(w).Encode(map[string]string{"statusCode": "200", "success": "true", "containerID": containerID, "stackID": id})
			}
		}
		if !found {
			json.NewEncoder(w).Encode(map[string]string{"statusCode": "400", "success": "false", "error": "an error occured"})
		}
	} else {
		json.NewEncoder(w).Encode(map[string]string{"statusCode": "400", "success": "false", "error": "an error occured"})
	}

}

func GET_dockercompose(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	result, err := t.ReadAllDockerComposeFiles()
	if err == nil {
		json.NewEncoder(w).Encode(map[string]interface{}{"statusCode": "200", "items": result})
	} else {
		json.NewEncoder(w).Encode(map[string]string{"statusCode": "400", "success": "false", "error": "an error occured"})
	}

}

func GET_ID_dockercompose(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]

	result, err := t.ReadDockerComposeFile(id)
	if err == nil {
		json.NewEncoder(w).Encode(map[string]interface{}{"statusCode": "200", "item": result})
	} else {
		json.NewEncoder(w).Encode(map[string]string{"statusCode": "400", "success": "false", "error": "item not found"})
	}

}
