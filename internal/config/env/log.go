package env

import (
	"errors"
	"os"
)

const (
	logDirPathEnvName = "LOG_DIR_PATH"
	envTypeName       = "ENV_TYPE"
)

// LogConfig contains the log directory path and environment type.
type LogConfig struct {
	logDirPath string // path to the log directory
	env        string // environment type (e.g. local, dev, prod)
}

// NewLogConfig creates a new LogConfig by getting the log directory path and environment type from the respective environment variables.
func NewLogConfig() (*LogConfig, error) {
	logDirPath := os.Getenv(logDirPathEnvName)
	if len(logDirPath) == 0 {
		return nil, errors.New("log dir path is not set")
	}

	env := os.Getenv(envTypeName)
	if len(env) == 0 {
		return nil, errors.New("env type is not set")
	}

	return &LogConfig{
		logDirPath: logDirPath,
		env:        env,
	}, nil
}

// DirPath returns the log directory path.
func (c *LogConfig) DirPath() string {
	return c.logDirPath
}

// Env returns the environment type.
func (c *LogConfig) Env() string {
	return c.env
}
