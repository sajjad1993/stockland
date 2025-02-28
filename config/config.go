// Package config provides types and functionalities for retrieving server configuration values
package config

import (
	"fmt"
	"github.com/spf13/viper"
	"time"
)

// DatabaseConfig represents database configuration values
type DatabaseConfig struct {
	DatabaseDsn     string
	DatabaseDisable bool
}

// HTTPServerConfig represents HTTP server configuration values
type HTTPServerConfig struct {
	HTTPServerReadTimeout       time.Duration
	HTTPServerReadHeaderTimeout time.Duration
	HTTPServerWriteTimeout      time.Duration
	HTTPServerAddress           string
	HTTPServerPort              uint
}

// JwtToken represents JWT token configuration values
type JwtToken struct {
	AccessSignKey  string
	AccessJWTExp   time.Duration
	RefreshSignKey string
	RefreshExp     time.Duration
}

// Mongo represents MongoDB connection details
type Mongo struct {
	MongoURI string
}

// V2rayAddress represents V2ray server addresses and ports
type V2rayAddress struct {
	MtnAddress   string
	MtnXAddress  string
	MciAddress   string
	MciXAddress  string
	WifiAddress  string
	WifiXAddress string
	Port443      uint
	Port80       uint
	Port2053     uint
	Port2083     uint
}

// Config represents the overall configuration
type Config struct {
	DatabaseConfig
	HTTPServerConfig
	GlobalTimeout time.Duration
	JwtToken
	Mongo
	V2rayAddress
	AIKey string
}

// NewConfigFromViper NewConfigFromEnv initializes configuration from environment variables and files
func NewConfigFromViper() Config {
	viper.SetConfigName("config") // Name of the config file (without extension)
	viper.SetConfigType("yaml")   // Config file type (YAML, JSON, TOML)
	viper.AddConfigPath(".")      // Look for config in the current directory
	viper.AutomaticEnv()          // Read values from environment variables

	// Set default values in case environment variables or config files are missing
	viper.SetDefault("DatabaseDsn", "host=0.0.0.0 user=sajjad password=sajjad123 dbname=postgres port=5435 TimeZone=UTC")
	viper.SetDefault("DatabaseDisable", true)
	viper.SetDefault("HTTPServerReadTimeout", "15s")
	viper.SetDefault("HTTPServerReadHeaderTimeout", "15s")
	viper.SetDefault("HTTPServerWriteTimeout", "15s")
	viper.SetDefault("HTTPServerAddress", "127.0.0.1:9010")
	viper.SetDefault("HTTPServerPort", 4444)
	viper.SetDefault("JwtToken.AccessSignKey", "hajihaji")
	viper.SetDefault("JwtToken.AccessJWTExp", "168h")
	viper.SetDefault("JwtToken.RefreshExp", "720h")
	viper.SetDefault("V2rayAddress.MtnAddress", "mtn.fullspeedy.online")
	viper.SetDefault("V2rayAddress.MtnXAddress", "mtnx.fullspeedy.online")
	viper.SetDefault("V2rayAddress.MciAddress", "mci.fullspeedy.online")
	viper.SetDefault("V2rayAddress.MciXAddress", "mcix.fullspeedy.online")
	viper.SetDefault("V2rayAddress.WifiAddress", "wifi.fullspeedy.online")
	viper.SetDefault("V2rayAddress.WifiXAddress", "wifix.fullspeedy.online")
	viper.SetDefault("V2rayAddress.Port443", 443)
	viper.SetDefault("V2rayAddress.Port80", 80)
	viper.SetDefault("V2rayAddress.Port2053", 2053)
	viper.SetDefault("V2rayAddress.Port2083", 2083)
	viper.SetDefault("GlobalTimeout", "5s")
	viper.SetDefault("AIkey", "5s")

	// Read configuration file (optional)
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("⚠️  Warning: Could not read config file: %v\n", err)
	}

	// Unmarshal configuration into the Config struct
	var cfg Config
	err := viper.Unmarshal(&cfg)
	if err != nil {
		fmt.Printf("Error : config went wrong ..  ")
		panic(err)
	}

	return cfg
}
