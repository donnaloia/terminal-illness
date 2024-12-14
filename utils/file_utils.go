package utils

import (
	"os"
	"path/filepath"
	"strings"
)

func GetConfigDir() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	configDir := filepath.Join(homeDir, ".terminal-illness")
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return "", err
	}

	return configDir, nil
}

func SaveURL(url string) error {
	configDir, err := GetConfigDir()
	if err != nil {
		return err
	}

	historyFile := filepath.Join(configDir, "url_history.txt")

	// Read existing URLs
	var urls []string
	if data, err := os.ReadFile(historyFile); err == nil {
		urls = strings.Split(strings.TrimSpace(string(data)), "\n")
	}

	// Add new URL at the beginning
	urls = append([]string{url}, urls...)

	// Keep only the most recent 20 URLs
	if len(urls) > 20 {
		urls = urls[:20]
	}

	// Write back to file
	return os.WriteFile(historyFile, []byte(strings.Join(urls, "\n")+"\n"), 0644)
}

func GetSavedURLs() []string {
	configDir, err := GetConfigDir()
	if err != nil {
		return []string{}
	}

	historyFile := filepath.Join(configDir, "url_history.txt")
	data, err := os.ReadFile(historyFile)
	if err != nil {
		return []string{}
	}

	return strings.Split(strings.TrimSpace(string(data)), "\n")
}
