package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	gcfg "gopkg.in/gcfg.v1"
)

type MainConfig struct {
	Server ServerConfig
}

type ServerConfig struct {
	Environment string
	Host        string
	Port        int
	Debug       bool
}

var (
	Main      MainConfig
	configDir = "./files/config/%s/main.ini"
)

func init() {
	env := os.Getenv("ENV")
	if env == "" {
		log.Println("No environment set. Using development")
		env = "development"
	}

	configDir = fmt.Sprintf(configDir, env)
}

func Init() error {
	log.Println("Load Config File")

	if err := gcfg.ReadFileInto(&Main, configDir); err != nil {
		return err
	}

	bytes, err := json.MarshalIndent(Main, "", "   ")
	if err != nil {
		return err
	}

	log.Printf("Config: \n%s\n", string(bytes))
	return nil
}
