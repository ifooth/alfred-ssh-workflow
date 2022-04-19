package ssh

import (
	"fmt"
	"strings"

	aw "github.com/deanishe/awgo"
)

func GetSliceString(rawScripts interface{}) []string {
	switch scripts := rawScripts.(type) {
	case string:
		return []string{scripts}
	case []string:
		return scripts
	case []interface{}:
		values := make([]string, len(scripts))
		for _, script := range scripts {
			if v, ok := script.(string); ok {
				values = append(values, v)
			}
		}
		return values
	default:
		return []string{}
	}
}

type SSH struct {
	Host        string      `yaml:"Host"`
	Hostname    string      `yaml:"Hostname"`
	Port        int         `yaml:"Port,omitempty"`
	User        string      `yaml:"User,omitempty"`
	Password    string      `yaml:"Password,omitempty"`
	PreScripts  interface{} `yaml:"PreScripts,omitempty"`
	PostScripts interface{} `yaml:"PostScripts,omitempty"`
}

func (s *SSH) String() string {
	return fmt.Sprintf("%s<%s>", s.Host, s.Hostname)
}

func (s *SSH) GetPreScripts() []string {
	return GetSliceString(s.PreScripts)
}

func (s *SSH) GetPostScripts() []string {
	return GetSliceString(s.PostScripts)
}

func (s *SSH) GetHost() string {
	user := s.User
	if user == "" {
		user = "root"
	}
	return fmt.Sprintf("%s@%s", user, s.Hostname)

}

// Arg 生成对应的query
func (s *SSH) GetArg() string {
	_query := []string{}
	preScripts := s.GetPreScripts()
	_query = append(_query, preScripts...)

	host := s.GetHost()
	sshCmd := fmt.Sprintf("ssh %s", host)
	if len(preScripts) > 0 {
		sshCmd = fmt.Sprintf("waitdone:ssh %s", s.Hostname)
	}

	if s.Port != 0 && s.Port != 22 {
		sshCmd += fmt.Sprintf(" -p %d", s.Port)
	}

	_query = append(_query, sshCmd)

	// 组装password命令
	password_cmd := fmt.Sprintf("password:%s", s.Password)
	_query = append(_query, password_cmd)

	postScripts := s.GetPostScripts()
	if len(postScripts) > 0 {
		_query = append(_query, postScripts...)
	}
	host_query := strings.Join(_query, "\n")

	return host_query
}

func (s *SSH) GetAutocomplete() string {
	return s.Host
}

func (s *SSH) AddItem(wf *aw.Workflow) {
	wf.NewItem(s.Host).
		Subtitle(s.GetHost()).
		Copytext(s.Hostname).
		Largetype(s.Hostname).
		Var("Hostname", s.Hostname). // 提供复制内容
		Icon(aw.IconWorkflow).Arg(s.GetArg()).
		Autocomplete(s.GetAutocomplete()).
		Valid(true)
}
