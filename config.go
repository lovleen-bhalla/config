package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"reflect"
)

var ErrKeyNotFound = errors.New("Key not in the config")

type Config struct {
	m map[string]any
}

func NewConfigFromFile(filename string) (*Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, errors.New(fmt.Sprintln("config: ", err.Error()))
	}
	return NewConfigFromReader(file)
}

func NewConfigFromReader(reader io.Reader) (*Config, error) {
	bytes, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, errors.New(fmt.Sprintln("config: ", err.Error()))
	}
	return NewConfig(bytes)
}

func NewConfig(data []byte) (*Config, error) {
	var m map[string]any
	err := json.Unmarshal(data, &m)
	if err != nil {
		return nil, errors.New(fmt.Sprintln("config: ", err.Error()))
	}
	return &Config{m}, nil
}

func (c *Config) get(key string) (any, error) {
	val, ok := c.m[key]
	if !ok {
		return nil, ErrKeyNotFound
	}
	return val, nil
}

func (c *Config) GetString(key string) (string, error) {
	val, err := c.get(key)
	if err != nil {
		return "", err
	}
	typ := reflect.TypeOf(val).String()
	if typ != "string" {
		return "", errors.New(fmt.Sprintf("config: value stored with key %s is of type %s", key, typ))
	}
	return val.(string), nil
}

func (c *Config) GetInt(key string) (int, error) {
	val, err := c.get(key)
	if err != nil {
		return 0, err
	}
	typ := reflect.TypeOf(val).String()
	if typ != "float64" {
		return 0, errors.New(fmt.Sprintf("config: value stored with key %s is of type %s", key, typ))
	}
	return int(val.(float64)), nil
}
