// Package config provides configuration struct for the program.
package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Options Options      `yaml:"options,omitempty"`
	Import  Import       `yaml:"import"`
	Export  []ExportItem `yaml:"export"`
}

type Options struct {
	DisplayValues bool `yaml:"displayValues"`
}

type Import struct {
	FromFile []string `yaml:"fromFile,omitempty"`
	FromURL  []string `yaml:"fromURL,omitempty"`
	FromText string   `yaml:"fromText,omitempty"`
}

type ExportItem struct {
	Template    string `yaml:"template"`
	Target      string `yaml:"target"`
	Description string `yaml:"description,omitempty"`
}

func Read(cfg *Config) error {
	// Configuration file
	viper.SetConfigName("dustomize") // name of config file (without extension)
	viper.SetConfigType("yaml")      // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")         // optionally look for config in the working directory

	// Read configuration
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	// Unmarshal the configuration
	if err := viper.Unmarshal(&cfg); err != nil {
		return err
	}
	return nil
}
