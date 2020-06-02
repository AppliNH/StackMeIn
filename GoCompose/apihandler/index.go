package apihandler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	t "primitivo.fr/applinh/go-docker-compose/dockercomposewrite"
	d "primitivo.fr/applinh/go-docker-compose/dockercontrol"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ok")
}

func GET_ResHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	result, err := t.ReadAllDockerComposeFiles()
	if err == nil {
		json.NewEncoder(w).Encode(map[string]interface{}{"statusCode": "200", "items": result})
	} else {
		json.NewEncoder(w).Encode(map[string]string{"statusCode": "400", "success": "false", "error": "an error occured"})
	}

}

func GET_Start(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	d.CreateNewContainer(id)
	d.ListContainers()

}

func GET_ID_ResHandler(w http.ResponseWriter, r *http.Request) {
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

func POST_ResHandler(w http.ResponseWriter, r *http.Request) {
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

// func PATCH_ResHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	res := mux.Vars(r)["res"]
// 	id := mux.Vars(r)["id"]

// 	data := make(map[string]interface{})
// 	err := json.NewDecoder(r.Body).Decode(&data)

// 	if err != nil {
// 		fmt.Println(err)
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	db.UpdateItem(res, id, data)
// 	response := map[string]string{"statusCode": "200", "success": "true", "id": id, "res": res}
// 	json.NewEncoder(w).Encode(response)
// }
