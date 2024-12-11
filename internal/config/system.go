package config

import (
	"os"
	"strconv"
)

const (
	debugEnvName = "DEBUG"
)

// SystemConfig represents a config that stores system settings.
type SystemConfig interface {
	Debug() bool
}

type systemConfig struct {
	debug bool
}

// NewSystemConfig creates a new SystemConfig.
func NewSystemConfig() (SystemConfig, error) {
	debug := os.Getenv(debugEnvName)

	// Skip error because we have a default value
	debugBool, _ := strconv.ParseBool(debug)
	return &systemConfig{
		debug: debugBool,
	}, nil
}

// Debug returns a debug mode.
func (s *systemConfig) Debug() bool {
	return s.debug
}
