package config

import (
	"encoding/json"
	"log"
	"os"
)

const (
	DefaultConfigPath = "~/.config/HyprWalls/config"
)

func FormatPath(path_to_file string) string {
	var temp_path string = path_to_file
	if temp_path[0] == '~' {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			log.Fatalf("failed to get home directory: %v", err)
		}
		temp_path = homeDir + temp_path[1:]
	}
	return temp_path
}

func LoadImagePath(path_to_config string) map[string]string {
	var file_content []byte
	var err error
	var jsonData []byte
	var data map[string]string

	if len(path_to_config) == 0 {
		path_to_config = DefaultConfigPath
	}

	path_to_config = FormatPath(path_to_config)

	file_content, err = os.ReadFile(path_to_config)
	if err != nil {
		log.Fatal(err)
	}

	jsonData = []byte{'{'}
	for i := 0; i < len(file_content); i++ {
		jsonData = append(jsonData, file_content[i])
	}
	jsonData = append(jsonData, '}')

	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		log.Fatal(err)
	}

	return data
}
