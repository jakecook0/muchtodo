package main

import (
	"errors"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Host string
	Port string
}

func configSetup() Config {
	// check that file exists
	var filePath string = "config.yaml"

	if !checkFileExists(filePath) {
		log.Fatalf("Config file %s not found, must be next to the server.", filePath)
	}

	configData, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	config := parseConfig(configData)

	log.Printf("Config: %v", config)

	return *config
}

func checkFileExists(filePath string) bool {
	_, error := os.Stat(filePath)
	return !errors.Is(error, os.ErrNotExist)
}

func parseConfig(configData []byte) *Config {
	config := Config{}

	err := yaml.Unmarshal(configData, &config)
	if err != nil {
		log.Fatalf("Error parsing YAML, %v", err)
	}

	return &config
}
