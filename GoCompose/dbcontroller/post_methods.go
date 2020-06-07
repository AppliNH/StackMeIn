package dbcontroller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func WriteToDB(ressource string, payload map[string]interface{}) (map[string]interface{}, error) {
	jsonpayload, _ := json.Marshal(payload)
	resp, err := http.Post("http://localhost:5000/"+ressource, "application/json", bytes.NewBuffer(jsonpayload))
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
