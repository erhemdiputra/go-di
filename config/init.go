package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
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
	Timeout     string
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
	configDir = "./files/config/%s/"
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
	viper.SetConfigType("toml")
	viper.SetConfigName("main")
	viper.AddConfigPath(configDir)

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if err := viper.Unmarshal(&globalCfg); err != nil {
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
