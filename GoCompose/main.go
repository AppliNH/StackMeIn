package main

import (
	"fmt"

	"log"
	"net/http"

	"github.com/gorilla/mux"
	api "primitivo.fr/applinh/go-docker-compose/apihandler"
)

func Rewriter(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.RequestURI)
	})
}

func main() {
	// services := make(map[string]interface{})
	// ubuntu := map[string]interface{}{"image": "ubuntu", "container_name": "ubuntu", "ports": []string{"5000:5000"}, "networks": []string{"main_network"}}
	// net := make(map[string]interface{})

	// net["main_network"] = ""
	// services["ubuntu"] = ubuntu

	//t.ParseComposeData("2.1", services, net)
	r := mux.NewRouter()
	r.HandleFunc("/stack/{id}", api.GET_Stack).Methods("GET")
	r.HandleFunc("/stack/{id}", api.DELETE_ID_stack).Methods("DELETE")

	r.HandleFunc("/dockercompose", api.GET_dockercompose).Methods("GET")         // TODO: DEPRECATED
	r.HandleFunc("/dockercompose/{id}", api.GET_ID_dockercompose).Methods("GET") // TODO: DEPRECATED
	r.HandleFunc("/dockercompose", api.POST_dockercompose).Methods("POST")
	// r.HandleFunc("/{res}/{id}", api.PATCH_ResHandler).Methods("PATCH")

	log.Fatal(http.ListenAndServe(":1997", r))

}
