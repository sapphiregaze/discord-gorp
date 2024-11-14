package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"

	"github.com/sapphiregaze/discord-gorp/pkg/logger"
)

type Timestamp struct {
	Start int64 `mapstructure:"start" json:"start,omitempty"`
	End   int64 `mapstructure:"end" json:"end,omitempty"`
}

type Assets struct {
	LargeImage string `mapstructure:"large_image" json:"large_image,omitempty"`
	LargeText  string `mapstructure:"large_text" json:"large_text,omitempty"`
	SmallImage string `mapstructure:"small_image" json:"small_image,omitempty"`
	SmallText  string `mapstructure:"small_text" json:"small_text,omitempty"`
}

type Party struct {
	ID   string `mapstructure:"id" json:"id,omitempty"`
	Size [2]int `mapstructure:"size" json:"size,omitempty"`
}

type Secrets struct {
	Join     string `mapstructure:"join" json:"join,omitempty"`
	Spectate string `mapstructure:"spectate" json:"spectate,omitempty"`
	Match    string `mapstructure:"match" json:"match,omitempty"`
}

type Button struct {
	Label string `mapstructure:"label" json:"label,omitempty"`
	Url   string `mapstructure:"url" json:"url,omitempty"`
}

type Activity struct {
	ApplicationID string     `mapstructure:"application_id" json:"application_id,omitempty"`
	Name          string     `mapstructure:"name" json:"name,omitempty"`
	Type          int        `mapstructure:"type" json:"type,omitempty"`
	State         string     `mapstructure:"state" json:"state,omitempty"`
	Details       string     `mapstructure:"details" json:"details,omitempty"`
	Timestamp     *Timestamp `mapstructure:"timestamp" json:"timestamps,omitempty"`
	Assets        *Assets    `mapstructure:"assets" json:"assets,omitempty"`
	Party         *Party     `mapstructure:"party" json:"party,omitempty"`
	Secrets       *Secrets   `mapstructure:"secrets" json:"secrets,omitempty"`
	Buttons       []Button   `mapstructure:"buttons" json:"buttons,omitempty"`
	Instance      bool       `mapstructure:"instance" json:"instance"`
}

type Config struct {
	Activity Activity `mapstructure:"activity"`
}

func copy(src string, dst string) error {
	data, err := os.ReadFile(src)
	if err != nil {
		return err
	}

	err = os.WriteFile(dst, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func Load() (*Config, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	configDir := filepath.Join(homeDir, ".config/discord-gorp")
	configPath := filepath.Join(configDir, "config.yaml")

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		if err := os.MkdirAll(configDir, 0755); err != nil {
			return nil, fmt.Errorf("failed to create config directory: %w", err)
		}

		currentDir, err := os.Getwd()
		if err != nil {
			return nil, err
		}

		logger.Info(fmt.Sprintf("Config file not found. Copying example config to: %s...", configPath))
		if err := copy(filepath.Join(currentDir, "configs/config.example.yaml"), configPath); err != nil {
			return nil, fmt.Errorf("failed to copy example config: %w", err)
		}
	}

	viper.SetConfigFile(configPath)
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
