package config

import "fmt"

func (c *Config) ValidateServer() error {
	if c.name != "srv" {
		return fmt.Errorf("app name expected 'srv' but is '%s'", c.name)
	}
	if c.GetString("server.host") == "" {
		return fmt.Errorf("property server.host missing")
	}
	if c.GetString("server.port") == "" {
		return fmt.Errorf("property server.port missing")
	}
	if c.GetString("bucket.name") == "" {
		return fmt.Errorf("property bucket.name missing")
	}
	return nil
}
