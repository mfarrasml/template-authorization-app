package config

import (
	"errors"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	dbUrl      string
	jwtSecret  string
	jwtIssuer  string
	jwtExpiry  int
	bcryptCost int
	host       string
	port       int
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, errors.New("error loading environment")
	}

	expiry, err := strconv.Atoi(os.Getenv("JWT_AUTH_TOKEN_EXPIRY_MINUTES"))
	if err != nil {
		return nil, errors.New("error loading jwt expiry env value")
	}

	cost, err := strconv.Atoi(os.Getenv("BCRYPT_COST"))
	if err != nil {
		return nil, errors.New("error loading bcrypt cost env value")
	}

	port, err := strconv.Atoi(os.Getenv("ROUTER_PORT"))
	if err != nil {
		return nil, errors.New("error loading port env value")
	}

	config := Config{
		dbUrl:      os.Getenv("DB_URL"),
		jwtSecret:  os.Getenv("JWT_SECRET"),
		jwtIssuer:  os.Getenv("ISSUER"),
		jwtExpiry:  expiry,
		bcryptCost: cost,
		host:       os.Getenv("ROUTER_HOST"),
		port:       port,
	}

	return &config, nil

}

func (c *Config) DbUrl() string {
	return c.dbUrl
}

func (c *Config) JwtSecret() string {
	return c.jwtSecret
}

func (c *Config) JwtIssuer() string {
	return c.jwtIssuer
}

func (c *Config) JwtExpiry() int {
	return c.jwtExpiry
}

func (c *Config) BcryptCost() int {
	return c.bcryptCost
}

func (c *Config) ApiHost() string {
	return c.host
}

func (c *Config) ApiPort() int {
	return c.port
}
