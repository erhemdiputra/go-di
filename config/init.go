package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	gcfg "gopkg.in/gcfg.v1"
)

type MainConfig struct {
	Server   ServerConfig
	Database DatabaseConfig
}

type ServerConfig struct {
	Environment string
	Host        string
	Port        int
	Debug       bool
}

type DatabaseConfig struct {
	Host     string
	Driver   string
	Username string
	Password string
	Name     string
}

var (
	globalCfg MainConfig
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
	if err := gcfg.ReadFileInto(&globalCfg, configDir); err != nil {
		return err
	}

	bytes, err := json.MarshalIndent(globalCfg, "", "   ")
	if err != nil {
		return err
	}

	log.Printf("Config loaded successfully: \n%s\n", string(bytes))
	return nil
}

func Get() MainConfig {
	return globalCfg
}

func (dc DatabaseConfig) String() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s)/%s",
		dc.Username, dc.Password, dc.Host, dc.Name,
	)
}
