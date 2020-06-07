package apihandler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	db "primitivo.fr/applinh/go-docker-compose/dbcontroller"
	t "primitivo.fr/applinh/go-docker-compose/dc_file_mng"
	d "primitivo.fr/applinh/go-docker-compose/dockercontrol"
)

func GET_Stack(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id := vars["id"]

	if res_getDB, err := db.Get_byid("stacks", id); err == nil {
		data := make(map[string]interface{})
		data = res_getDB["dockercompose"].(map[string]interface{})
		if _, erro := t.ParseComposeData(id, data["version"].(string), data["services"].(map[string]interface{}), data["networks"].(map[string]interface{})); erro == nil {
			if res, err1 := d.CreateNewContainer(id); err1 == nil {
				containers, _ := d.ListContainers()
				found := false
				for _, v := range containers {
					if v.ID == res["containerID"] {
						found = true
						res_getDB["containers"] = res["otherContainers"]
						db.Patch_byid("stacks", id, res_getDB)
						d.StopContainer(res["containerID"].(string))
						json.NewEncoder(w).Encode(map[string]string{"statusCode": "200", "stackID": id})
					}
				}
				if !found {
					json.NewEncoder(w).Encode(map[string]string{"statusCode": "400", "error": "an error occured"})
				}
			} else {
				json.NewEncoder(w).Encode(map[string]string{"statusCode": "400", "error": "an error occured"})
			}
		}
	}

}

func GET_dockercompose(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	result, err := t.ReadAllDockerComposeFiles()
	if err == nil {
		json.NewEncoder(w).Encode(map[string]interface{}{"statusCode": "200", "items": result})
	} else {
		json.NewEncoder(w).Encode(map[string]string{"statusCode": "400", "error": "an error occured"})
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
		json.NewEncoder(w).Encode(map[string]string{"statusCode": "400", "error": "item not found"})
	}

}
