package config

import (
	"fmt"
	"os"
	"strings"
)

func loadSecret(s string, key string) (string, error) {
	if s == "" {
		return "", fmt.Errorf("missing parameter %s", key)
	}

	if !strings.HasPrefix(s, "file://") {
		return s, nil
	}

	data, err := os.ReadFile(strings.TrimPrefix(s, "file://"))
	if err != nil {
		return "", err
	}

	val := strings.TrimSpace(string(data))
	if val == "" {
		return "", fmt.Errorf("file is empty %s", key)
	}

	return val, nil
}
