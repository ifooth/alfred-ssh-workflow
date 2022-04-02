package config

import (
	"os"

	aw "github.com/deanishe/awgo"
	"github.com/ifooth/alfred-ssh-workflow/ssh"
	"gopkg.in/yaml.v3"
)

type SSHConfig struct {
	Provider string `yaml:"provider"`
	Path     string `yaml:"path"`
}

func (c *SSHConfig) HandleItem(wf *aw.Workflow) error {
	data, err := os.ReadFile(c.Path)
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
