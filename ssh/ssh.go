package ssh

type SSH struct {
	Host        string   `yaml:"host"`
	Hostname    string   `yaml:"hostname"`
	User        string   `yaml:"user,omitempty"`
	Password    string   `yaml:"password,omitempty"`
	PreScripts  []string `yaml:"pre_scripts,omitempty"`
	PostScripts []string `yaml:"post_scripts,omitempty"`
}
