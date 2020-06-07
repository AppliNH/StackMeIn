package apihandler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	db "primitivo.fr/applinh/go-docker-compose/dbcontroller"
	d "primitivo.fr/applinh/go-docker-compose/dockercontrol"
)

func DELETE_ID_stack(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)

	stackId := vars["id"]

	res_get, err := db.Get_byid("stacks", stackId)
	if err == nil && len(res_get["containers"].([]interface{})) > 0 {
		containers := res_get["containers"].([]interface{})
		remaining_containers := containers
		for k, containerId := range containers {
			if err := d.StopContainer(containerId.(string)); err != nil {
				json.NewEncoder(w).Encode(map[string]string{"statusCode": "400", "error": "an error occured", "containerID": containerId.(string)})
			} else {
				if k < len(remaining_containers) {
					remaining_containers = append(remaining_containers[:k], remaining_containers[k+1:]...)
				} else {
					remaining_containers = []interface{}{}
				}
				json.NewEncoder(w).Encode(map[string]string{"statusCode": "200", "message": "stopping container" + containerId.(string)})
			}
		}
		res_get["containers"] = remaining_containers
		db.Patch_byid("stacks", stackId, res_get)
	}
}
