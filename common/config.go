package common

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

const (
	Host       = "http://localhost:8081"
	ConfigFile = "/etc/service/service.json"
)

type Config struct {
	User   string `json:"user"`
	Host   string `json:"host"`
	Loaded bool   `json:"-"`
}

func readConfig() []byte {
	data, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func writeConfig(c *Config) error {
	str, _ := json.Marshal(c)
	err := ioutil.WriteFile(ConfigFile, []byte(str), 0777)
	if err != nil {
		log.Fatal(err)
	}

	return err
}

// SaveConfig is the interface to writing to the config file.
func (c *Config) SaveConfig(configFile string) error {
	err := writeConfig(c)
	if err != nil {
		return err
	}
	return nil
}

// LoadConfig attempts to load configuration. Will create a default config file if one doesn't exist.
func (c *Config) LoadConfig() error {
	// attempt to see if config file is present
	_, err := os.Stat(ConfigFile)

	if os.IsNotExist(err) {
		fmt.Print("\n")
		fmt.Println("--------")
		fmt.Println("No configuration file found. Creating one now.")
		fmt.Println("--------")
		fmt.Print("\n")
		c.Host = "http://localhost:8081"
		c.User = ""
		c.Loaded = true
		writeConfig(c)
		return nil
	} else if err != nil {
		return err
	}

	return nil
}

func NewConfig() *Config {
	return &Config{
		User: "placeholder",
		Host: "http://localhost",
	}
}

// SetUsername sets username to use for tasker.
func (c *Config) SetUsername(username string) error {
	c.User = username
	err := writeConfig(c)
	return err
}

// GetUsername returns the username as found in the config.
func (c *Config) GetUsername() string {
	data := readConfig()
	json.Unmarshal([]byte(data), &c)
	return c.User
}

// SetHost sets host adderess to use for tasker.
func (c *Config) SetHost(host string) error {
	c.Host = host
	err := writeConfig(c)
	return err
}

// GetHost returns the host as found in the config.
func (c *Config) GetHost() string {
	data := readConfig()
	json.Unmarshal([]byte(data), &c)
	return c.Host
}
