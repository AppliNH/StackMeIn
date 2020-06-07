package dbcontroller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func Get_byid(ressource string, id string) (map[string]interface{}, error) {
	resp, err := http.Get("http://localhost:5000/" + ressource + "/" + id)

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
