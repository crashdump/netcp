package config

import "fmt"

func (c *Config) ValidateClient() error {
	if c.name != "cli" {
		return fmt.Errorf("wrong app name, is '%s' expected 'cli", c.name)
	}
	if c.GetString("server.host") == "" {
		return fmt.Errorf("property server.host missing")
	}
	if c.GetString("server.port") == "" {
		return fmt.Errorf("property server.port missing")
	}
	return nil
}
