package dbcontroller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func Get_byid(ressource string, id string) (map[string]interface{}, error) {
	var url string
	mode := os.Getenv("MODE")

	if mode == "COMPOSE" {
		fmt.Println("yes")
		url = "http://firego:5000/"
	} else {
		url = "http://localhost:5000/"
	}

	resp, err := http.Get(url + ressource + "/" + id)

	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(body))
	var res map[string]interface{}
	if err := json.Unmarshal(body, &res); err != nil {
		panic(err)
	}

	return res, nil

}
