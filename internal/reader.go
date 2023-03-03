// Package internal provides internal helper functions.
package internal

import (
	"io/ioutil"
	"net/http"
	"os"

	"gopkg.in/yaml.v3"
)

// Read & export directly from YAML string
func ReadConfigFromFile(config string) (map[interface{}]interface{}, error) {
	fileContent, err := os.ReadFile(config)
	if err != nil {
		return nil, err
	}

	mapResult := make(map[interface{}]interface{})
	mapResult, err = ReadConfigFromText(string(fileContent))
	if err != nil {
		return nil, err
	}

	return mapResult, nil
}

// Read & export from a YAML file
func ReadConfigFromText(config string) (map[interface{}]interface{}, error) {
	mapResult := make(map[interface{}]interface{})
	err := yaml.Unmarshal([]byte(config), &mapResult)

	if err != nil {
		return nil, err
	}

	return mapResult, nil
}

// Read & export from a YAML file in the web
func ReadConfigFromUrl(url string) (map[interface{}]interface{}, error) {
	response, err := http.Get(url)

	if err != nil {
		StopWithDebug(ErrBadURLRequest, err)
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		Stop(ErrBadURLHTTPResposne)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		StopWithDebug(ErrBadURLResponse, err)
	}

	// print response body
	resp := string(body)

	mapResult := make(map[interface{}]interface{})
	err = yaml.Unmarshal([]byte(resp), &mapResult)

	if err != nil {
		return nil, err
	}

	return mapResult, nil
}
