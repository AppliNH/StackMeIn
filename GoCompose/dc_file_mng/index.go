package dc_file_mng

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/mitchellh/mapstructure"
	"gopkg.in/yaml.v2"
	m "primitivo.fr/applinh/go-docker-compose/models"
	utils "primitivo.fr/applinh/go-docker-compose/utils"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func createServices(data map[string]interface{}) map[string]m.Service {
	services := make(map[string]m.Service)

	for k, v := range data {
		var result m.Service
		mapstructure.Decode(v, &result)
		services[k] = result
	}

	return services
}

func createNetworks(data map[string]interface{}) map[string]m.Network {
	networks := make(map[string]m.Network)

	for k, v := range data {
		var result m.Network
		mapstructure.Decode(v, &result)
		networks[k] = result
	}

	return networks
}

func ParseComposeData(uuid string, version string, services map[string]interface{}, networks map[string]interface{}) (string, error) {

	servicesObj := createServices(services)
	networksObj := createNetworks(networks)

	return InitDockerComposeFile(uuid, version, servicesObj, networksObj)

}

func InitDockerComposeFile(uuid string, version string, services map[string]m.Service, networks map[string]m.Network) (string, error) {

	t := m.T{Version: version, Services: services, Networks: networks}

	d, err := yaml.Marshal(&t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	if uuid, erro := utils.WriteDockerComposeFile(uuid, string(d)); erro != nil {
		return "", erro
	} else {
		return uuid, nil
	}
}

func ReadDockerComposeFile(uuid string) (m.T, error) {

	dat, err := ioutil.ReadFile("./composefiles/" + uuid + "/docker-compose.yml")
	if err != nil {
		fmt.Println(err)
	}
	t := m.T{}
	erro := yaml.Unmarshal([]byte(dat), &t)
	if erro != nil {
		fmt.Println(err)
	}
	return t, err

}
func ReadAllDockerComposeFiles() ([]m.T, error) {

	files, err := ioutil.ReadDir("./composefiles")
	if err != nil {
		fmt.Println(err)
	}

	DCfiles := []m.T{}

	for _, f := range files {
		if res, err := ReadDockerComposeFile(f.Name()); err == nil {
			DCfiles = append(DCfiles, res)
		}
		fmt.Println(f.Name())
	}

	return DCfiles, err

}
