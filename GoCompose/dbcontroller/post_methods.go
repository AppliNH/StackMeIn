package dbcontroller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func WriteToDB(ressource string, payload map[string]interface{}) (map[string]interface{}, error) {

	var url string
	mode := os.Getenv("MODE")

	fmt.Print(mode)

	if mode == "COMPOSE" {
		fmt.Println("yes")
		url = "http://firego:5000/"
	} else {
		url = "http://localhost:5000/"
	}

	jsonpayload, _ := json.Marshal(payload)
	resp, err := http.Post(url+ressource, "application/json", bytes.NewBuffer(jsonpayload))
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
