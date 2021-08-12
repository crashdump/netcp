package config

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"os"
	"os/user"
)

type Config struct {
	viper    viper.Viper
	name     string
	env      string
	version  string
	defaults map[string]interface{}
}

func New(name string, env string, defaults map[string]interface{}) (*Config, error) {
	if name == "" || env == "" {
		return nil, errors.New("please provide 'name' and 'env'")
	}

	user, err := user.Current()
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf("%s/.netcp/", user.HomeDir)

	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0700)
	}

	v := viper.New()
	v.SetConfigName("netcp" + name + "." + env)
	v.AddConfigPath(path)

	for key, value := range defaults {
		v.SetDefault(key, value)
	}
	//v.SetDefault("", defaults)

	return &Config{
		viper:    *v,
		name:     name,
		env:      env,
		defaults: defaults,
	}, nil
}

func (c *Config) Load() error {
	fmt.Printf("Loading %s\n", c.viper.ConfigFileUsed())

	err := c.viper.ReadInConfig()
	if err != nil {
		fmt.Print("error reading config file: ", err)
		return err
	}
	return nil
}

func (c *Config) Save() error {
	err := c.viper.WriteConfig()
	if err != nil {
		return err
	}
	return nil
}

func (c *Config) GetString(key string) string {
	return c.viper.GetString(key)
}

func (c *Config) GetInt(key string) int {
	return c.viper.GetInt(key)
}

func (c *Config) GetBool(key string) bool {
	return c.viper.GetBool(key)
}

func (c *Config) Set(key string, value interface{}) {
	c.viper.Set(key, value)
}

func (c *Config) GetName() string {
	return c.name
}

func (c *Config) GetEnv() string {
	return c.env
}

func (c *Config) GetVersion() string {
	return c.version
}
