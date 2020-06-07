package apihandler

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
