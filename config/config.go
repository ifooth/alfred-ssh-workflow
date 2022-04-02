package config

import (
	"os"
	"path/filepath"

	aw "github.com/deanishe/awgo"
	homedir "github.com/mitchellh/go-homedir"
	"gopkg.in/yaml.v3"

	"github.com/ifooth/alfred-ssh-workflow/ssh"
)

func AbsPath(path string) (string, error) {
	path, err := homedir.Expand(path)
	if err != nil {
		return "", err
	}
	return filepath.Abs(path)
}

type SSHConfig struct {
	Provider string `yaml:"provider"`
	Path     string `yaml:"path"`
}

func (c *SSHConfig) HandleItem(wf *aw.Workflow) error {
	path, err := AbsPath(c.Path)
	if err != nil {
		return err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	conf := []*ssh.SSH{}
	if err := yaml.Unmarshal(data, &conf); err != nil {
		return err
	}

	for _, c := range conf {
		c.AddItem(wf)
	}
	return nil
}

type Config struct {
	SSHConfigs []*SSHConfig `yaml:"data"`
}

func ReadConfig(path string) (*Config, error) {
	path, err := AbsPath(path)
	if err != nil {
		return nil, err
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	conf := new(Config)
	if err := yaml.Unmarshal(data, conf); err != nil {
		return nil, err
	}
	return conf, nil
}
