package ssh

import (
	"fmt"

	aw "github.com/deanishe/awgo"
)

type SSH struct {
	Host        string   `yaml:"Host"`
	Hostname    string   `yaml:"Hostname"`
	User        string   `yaml:"User,omitempty"`
	Password    string   `yaml:"Password,omitempty"`
	PreScripts  []string `yaml:"PreScripts,omitempty"`
	PostScripts []string `yaml:"PostScripts,omitempty"`
}

func (s *SSH) String() string {
	return fmt.Sprintf("%s<%s>", s.Host, s.Hostname)
}

// Arg 生成对应的query
func (s *SSH) GetArg() string {
	return ""
}

func (s *SSH) GetAutocomplete() string {
	return ""
}

func (s *SSH) AddItem(wf *aw.Workflow) {
	wf.NewItem(s.Host).
		Copytext(s.Hostname).
		Largetype(s.Hostname).
		Icon(aw.IconBurn).Arg(s.GetArg()).
		Autocomplete(s.GetAutocomplete())
}
