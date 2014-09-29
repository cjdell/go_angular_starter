package config

import (
	"fmt"
	"github.com/kylelemons/go-gypsy/yaml"
	"os"
	"path"
)

var App Config = &config{} // Main config instance

type (
	Config interface {
		Env() string
		WebRoot() string
		DatabaseConfig() (string, string, error)
	}

	config struct{}
)

func (config) Env() string {
	env := os.Getenv("ENV")

	if env == "" {
		env = "development"
	}

	return env
}

func (config) WebRoot() string {
	return "/home/cjdell/gopath/src/github.com/cjdell/go_angular_starter/web"
}

// Attempts to read database config file in "db/dbconf.yml"
func (self config) DatabaseConfig() (string, string, error) {
	var (
		cfgFile = path.Join("db", "dbconf.yml")
		env     = self.Env()
	)

	f, err := yaml.ReadFile(cfgFile)
	if err != nil {
		return "", "", err
	}

	drv, err := f.Get(fmt.Sprintf("%s.driver", env))
	if err != nil {
		return "", "", err
	}
	drv = os.ExpandEnv(drv)

	open, err := f.Get(fmt.Sprintf("%s.open", env))
	if err != nil {
		return "", "", err
	}
	open = os.ExpandEnv(open)

	return drv, open, nil
}
