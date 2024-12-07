package env

import (
	"errors"
	"fmt"
	"os"
)

const (
	pgHostEnvName     = "POSTGRES_HOST"
	pgPortEnvName     = "POSTGRES_PORT"
	pgDBNameEnvName   = "POSTGRES_DB"
	pgUserEnvName     = "POSTGRES_USER"
	pgPasswordEnvName = "POSTGRES_PASSWORD"
)

// PGConfig contains the PostgreSQL connection configuration details.
type PGConfig struct {
	host     string // the database server host
	port     string // the port on which the database server is listening
	name     string // the name of the database
	user     string // the username for authentication
	password string // the password for authentication
}

// NewPGConfig creates a new PGConfig by getting the PostgreSQL connection
// configuration details from the respective environment variables.
// If any of the environment variables is not set, it returns an error.
func NewPGConfig() (*PGConfig, error) {
	// Get the PostgreSQL database server host
	host := os.Getenv(pgHostEnvName)
	if len(host) == 0 {
		return nil, errors.New("postgres host is not set")
	}

	// Get the port on which the PostgreSQL database server is listening
	port := os.Getenv(pgPortEnvName)
	if len(port) == 0 {
		return nil, errors.New("postgres port is not set")
	}

	// Get the name of the PostgreSQL database
	name := os.Getenv(pgDBNameEnvName)
	if len(name) == 0 {
		return nil, errors.New("postgres db name is not set")
	}

	// Get the username for authentication
	user := os.Getenv(pgUserEnvName)
	if len(user) == 0 {
		return nil, errors.New("postgres user is not set")
	}

	// Get the password for authentication
	password := os.Getenv(pgPasswordEnvName)
	if len(password) == 0 {
		return nil, errors.New("postgres password is not set")
	}

	// Create a new PGConfig and return it
	return &PGConfig{
		host:     host,
		port:     port,
		name:     name,
		user:     user,
		password: password,
	}, nil
}

// DSN returns the Data Source Name (DSN) for the PostgreSQL connection.
//
// The DSN is in the following format:
//
//	host=<host> port=<port> user=<user> password=<password> dbname=<name> sslmode=disable
func (c *PGConfig) DSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", c.host, c.port, c.user, c.password, c.name)
}
