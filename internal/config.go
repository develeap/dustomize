package internal

import (
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
