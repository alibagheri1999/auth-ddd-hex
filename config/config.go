package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const (
	// ProductionMode indicates mode is release.
	ProductionMode = "production"
	// DevelopingMode indicates mode is debug.
	DevelopingMode = "developing"
)

var cfg *Config

// Init set the first and necessary settings
func InitAppConfig() {
	cfg = new(Config)
	v := viper.New()
	v.SetConfigType("yaml")
	v.SetConfigName("config")
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal("Error getting current directory:", err)
	}
	rootDir := findRootDirectory(currentDir, "go.mod")
	if rootDir == "" {
		log.Fatal("Root directory not found.")
	}
	// Set the configuration path to the project root directory
	v.AddConfigPath(filepath.Join(rootDir, "config"))
	if err1 := v.ReadInConfig(); err != nil {
		log.Fatal("error loading default configs: ", err1)
	}

	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	v.AutomaticEnv()
	if err := v.MergeInConfig(); err != nil {
		log.Println("no config file found. Using defaults and environment variables.")
	}
	if err := v.UnmarshalExact(&cfg); err != nil {
		log.Fatalf("invalid config schema: %v", err)
	}
}

func Init(configType string) {
	cfg = new(Config)
	v := viper.New()
	v.SetConfigType("yaml")
	if configType != "" {
		v.SetConfigName(configType + ".config")
	} else {
		v.SetConfigName("config")
	}
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal("Error getting current directory:", err)
	}
	rootDir := findRootDirectory(currentDir, "go.mod")
	if rootDir == "" {
		log.Fatal("Root directory not found.")
	}
	// Set the configuration path to the project root directory
	v.AddConfigPath(filepath.Join(rootDir, "config"))
	if err1 := v.ReadInConfig(); err != nil {
		log.Fatal("error loading default configs: ", err1)
	}

	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	v.AutomaticEnv()
	if err := v.MergeInConfig(); err != nil {
		log.Println("no config file found. Using defaults and environment variables.")
	}
	if err := v.UnmarshalExact(&cfg); err != nil {
		log.Fatalf("invalid config schema: %v", err)
	}
}

// Set ,set new configs
func Set(newConfig *Config) {
	cfg = newConfig
}

// Get ,get configs
func Get() *Config {
	return cfg
}

// IsProduction change to production mode
func (c *Config) IsProduction() bool {
	return c.App.Env == ProductionMode
}

// IsDebugging change to debugging mode
func (c *Config) IsDeveloping() bool {
	return c.App.Env == DevelopingMode
}

// findRootDirectory find the root of project
func findRootDirectory(currentDir, markerFile string) string {
	// Traverse up the directory tree looking for the marker file
	for {
		if _, err := os.Stat(filepath.Join(currentDir, markerFile)); err == nil {
			return currentDir
		}

		// Move up one directory
		newDir := filepath.Dir(currentDir)
		if newDir == currentDir {
			break
		}
		currentDir = newDir
	}
	return ""
}
