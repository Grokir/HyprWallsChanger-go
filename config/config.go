package config

import (
	"log"
	"os"
)

func FormatPath(path_to_file string) string {
	var temp_path string = path_to_file
	if temp_path[0] == '~' {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			log.Fatalf("failed to get home directory: %w", err)
		}
		temp_path = homeDir + temp_path[1:]
	}
	return temp_path
}
