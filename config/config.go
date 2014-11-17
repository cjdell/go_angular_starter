package config

import (
	"fmt"
	"github.com/kylelemons/go-gypsy/yaml"
	"os"
	"path"
)

var App *Config = newConfig() // Main config instance

type Config struct {
	DatabaseDriver string
	DatabaseOpen   string

	AssetRoot    string
	AdminRoot    string
	TemplateRoot string

	ListenAddress string
}

func newConfig() *Config {
	config := &Config{}
	config.ReadDatabaseConfig()
	config.ReadGeneralConfig()
	return config
}

func (Config) Env() string {
	env := os.Getenv("ENV")

	if env == "" {
		env = "development"
	}

	return env
}

// Attempts to read database config file in "db/dbconf.yml"
func (self *Config) ReadDatabaseConfig() error {
	var (
		cfgFile = path.Join("db", "dbconf.yml")
		f       *yaml.File
		err     error
	)

	if f, err = yaml.ReadFile(cfgFile); err != nil {
		return err
	}

	if self.DatabaseDriver, err = self.getParam(f, "driver"); err != nil {
		return err
	}

	if self.DatabaseOpen, err = self.getParam(f, "open"); err != nil {
		return err
	}

	return nil
}

// Attempts to read general config file in "config/config.yml"
func (self *Config) ReadGeneralConfig() error {
	var (
		cfgFile = path.Join("config", "config.yml")
		f       *yaml.File
		err     error
	)

	if f, err = yaml.ReadFile(cfgFile); err != nil {
		return err
	}

	if self.AssetRoot, err = self.getParam(f, "asset_root"); err != nil {
		return err
	}

	if self.AdminRoot, err = self.getParam(f, "admin_root"); err != nil {
		return err
	}

	if self.TemplateRoot, err = self.getParam(f, "template_root"); err != nil {
		return err
	}

	if self.ListenAddress, err = self.getParam(f, "listen_address"); err != nil {
		return err
	}

	return nil
}

func (self *Config) getParam(f *yaml.File, param string) (string, error) {
	env := self.Env()

	val, _ := f.Get(fmt.Sprintf("%s.%s", env, param))

	if val != "" {
		return val, nil
	}

	return f.Get(fmt.Sprintf("all.%s", param))
}
