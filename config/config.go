package config

import (
	"ai-smart-audit/internal/core/constants"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
)

type Config struct {
	AppConfig
}

type Provider struct {
	Config    *koanf.Koanf
	YamlFiles []string
	JsonFiles []string
}

func NewAppConfig() Config {
	configPath := os.Getenv(constants.ConfigPath)
	fmt.Println("Loading configuration from:", configPath)
	if configPath == "" {
		configPath = "./config/config.yaml" // Default path if not set
	}

	var ConfigFile = flag.String("configPath", configPath, "Path to the configuration file")
	flag.Parse()
	if *ConfigFile == "" {
		log.Fatal("Unable to load environment config")
	}

	var configFiles = []string{*ConfigFile}
	var provider = getProvider(configFiles...)
	return initAppConfig(provider)
}

func getProvider(files ...string) *Provider {
	return &Provider{
		Config:    koanf.New("."),
		YamlFiles: files,
		JsonFiles: nil,
	}
}

func initAppConfig(ko *Provider) Config {
	ko.Yaml()
	ko.Env()
	var appConfig Config
	err := ko.Config.UnmarshalWithConf("", &appConfig, koanf.UnmarshalConf{Tag: "json"})
	if err != nil {
		log.Fatalf("error unmarshalling config: %v", err)
	}
	return appConfig
}

func (p *Provider) Yaml() {
	for _, f := range p.YamlFiles {
		if err := p.Config.Load(file.Provider(f), yaml.Parser()); err != nil {
			if os.IsNotExist(err) {
				log.Fatalf("config file not found. please pass a configuration file. %v", err)
			}
			log.Fatalf("error loading config from file: %v", err)
		}
	}
}

func (p *Provider) Env() {
	if err := p.Config.Load(env.Provider("ENV_", ".", func(s string) string {
		return strings.ReplaceAll(strings.ToLower(
			strings.TrimPrefix(s, "ENV_")), "_", ".")
	}), nil); err != nil {
		log.Fatalf("error loading config from env: %v", err)
	}
}

func (c Config) Validate(validators ...func() error) {
	var err error
	for _, validator := range validators {
		if err = validator(); err != nil {
			log.Fatalf("validation error: %v", err)
		}
	}
}
