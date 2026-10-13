package config

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path"

	_ "embed"

	"github.com/IridiumNan/decision-analyst/internal/models"
	"github.com/spf13/viper"
)

//go:embed default.toml
var defaultConf []byte

var searchDirs = []string{
	".",
	path.Join(models.HomeDir(), ".config", models.AppName),
}

const configFileName = "config.toml"

func init() {
	loadConfig()
}

// loadConfig load configuration from [searchDirs] using viper
func loadConfig() {
	viper.SetConfigFile(configFileName)

	// viper.SetOptions(viper.KeyDelimiter("::"))

	for _, dir := range searchDirs {
		viper.AddConfigPath(dir)
	}

	err := viper.ReadInConfig()
	if err != nil {
		err := initConfig()
		if err != nil {
			log.Fatalf("error while init config, exiting..., err: %s", err.Error())
		}
	}
}

// initConfig write a new configuration toml file with comment
func initConfig() error {
	err := viper.ReadConfig(bytes.NewBuffer(defaultConf))
	if err != nil {
		return fmt.Errorf("error while unmarshal default configuration, update your code !!! err: %s", err.Error())
	}

	err = os.WriteFile(configFileName, defaultConf, 0o644)
	if err != nil {
		return fmt.Errorf("error while writing default configuration to file, err: %s", err.Error())
	}

	return nil
}
