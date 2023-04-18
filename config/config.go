package config

import (
	"encoding/json"
	"os"

	"github.com/labstack/gommon/log"
)

type Config struct {
	SPCliPath     string   `json:"sp_cli_path"`
	SPCliPort     int      `json:"sp_cli_port"`
	StorageDBPath string   `json:"storage_db_path"`
	Port          int      `json:"port"`
	WhitelistIPs  []string `json:"whitelist_ip"`
}

var StorageConfig Config = Config{
	SPCliPath:     "/home/ton-build/storage/storage-daemon/storage-daemon-cli",
	SPCliPort:     5555,
	StorageDBPath: "/home/ton-build/storage-db",
	Port:          34312,
	WhitelistIPs:  []string{""},
}

func LoadConfig() {
	file, err := os.Open("config.json")
	if err != nil {
		log.Info("Config file not found, using default config")
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	fileConfig := Config{}
	err = decoder.Decode(&fileConfig)
	if err != nil {
		log.Info("Config file is empty, using default config")
		return
	}

	// if config empty write default config
	if fileConfig.SPCliPath != "" {
		StorageConfig.SPCliPath = fileConfig.SPCliPath
	}
	if fileConfig.SPCliPort != 0 {
		StorageConfig.SPCliPort = fileConfig.SPCliPort
	}
	if fileConfig.StorageDBPath != "" {
		StorageConfig.StorageDBPath = fileConfig.StorageDBPath
	}
	if fileConfig.Port != 0 {
		StorageConfig.Port = fileConfig.Port
	}
	if len(fileConfig.WhitelistIPs) > 0 {
		StorageConfig.WhitelistIPs = fileConfig.WhitelistIPs
	}
}
