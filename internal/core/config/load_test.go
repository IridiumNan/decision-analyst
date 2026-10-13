package config

import (
	"bytes"
	"os"
	"testing"

	"github.com/spf13/viper"
)

var checker *testing.T

func setChecker(t *testing.T) {
	checker = t
}

func check[T comparable](want, got T) {
	if got != want {
		checker.Errorf("not match, want: %v, got: %v", want, got)
	}
}

func TestConfigLoad(t *testing.T) {
	setChecker(t)

	check(".", viper.GetString("database.dir"))

	check("127.0.0.1", viper.GetString("serve.addr"))
	check("8080", viper.GetString("serve.port"))

	err := os.Remove(configFileName)
	if err != nil {
		t.Errorf("error while removing legacy configuraion file, err: %s", err)
	}
}

// TestWithModify: modify defaultConf then write to file
// then load from file
func TestWithModify(t *testing.T) {
	setChecker(t)

	// replace 127.0.0.1 with 192.168.1.1
	modifiedConf := defaultConf

	expAddr := "192.168.1.1"
	expPort := "9090"
	modifiedConf = bytes.ReplaceAll(modifiedConf, []byte("127.0.0.1"), []byte(expAddr))
	modifiedConf = bytes.ReplaceAll(modifiedConf, []byte("8080"), []byte(expPort))

	err := os.WriteFile(configFileName, modifiedConf, 0o644)
	if err != nil {
		t.Errorf("error while writing modified configuration file, err: %s", err.Error())
	}

	loadConfig()

	check(".", viper.GetString("database.dir"))

	check(expAddr, viper.GetString("serve.addr"))
	check(expPort, viper.GetString("serve.port"))

	err = os.Remove(configFileName)
	if err != nil {
		t.Errorf("error while removing legacy configuraion file, err: %s", err)
	}
}
