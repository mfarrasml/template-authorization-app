package config

import (
	"errors"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	dbUrl           string
	jwtSecret       string
	jwtIssuer       string
	jwtAccTknExpiry int
	jwtRefTknExpiry int
	bcryptCost      int
	host            string
	port            int
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, errors.New("error loading environment")
	}

	accExp, err := strconv.Atoi(os.Getenv("JWT_ACCESS_TOKEN_EXPIRY_MINUTES"))
	if err != nil {
		return nil, errors.New("error loading jwt access token expiry env value")
	}

	refExp, err := strconv.Atoi(os.Getenv("JWT_REFRESH_TOKEN_EXPIRY_MINUTES"))
	if err != nil {
		return nil, errors.New("error loading jwt refresh token expiry env value")
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
		dbUrl:           os.Getenv("DB_URL"),
		jwtSecret:       os.Getenv("JWT_SECRET"),
		jwtIssuer:       os.Getenv("ISSUER"),
		jwtAccTknExpiry: accExp,
		jwtRefTknExpiry: refExp,
		bcryptCost:      cost,
		host:            os.Getenv("ROUTER_HOST"),
		port:            port,
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

func (c *Config) JwtAccTknExpiry() int {
	return c.jwtAccTknExpiry
}

func (c *Config) JwtRefTknExpiry() int {
	return c.jwtRefTknExpiry
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
