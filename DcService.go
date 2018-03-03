package ktdc

import (
	"strings"
	"errors"
)

type DcService struct {
	Image string                `json:"image"`
	ContainerName string        `json:"container_name,omitempty"`
	Ports []string              `json:"ports,omitempty"`
	Environment []UnixEnvString `json:"environment,omitempty"`
	Logging Logging `json:"logging,omitempty"`
	Command string `json:"command,omitempty"`
}

func (s *DcService) addEnv(key string, value string)  {
	s.Environment = append(s.Environment, UnixEnvString{key,value})
}

func (s *DcService) overwriteEnv(key string, value string)  {
	for k, v := range s.Environment {
		if v.Key == key {
			s.Environment[k].Value = value
			return
		}
	}
}

func (s *DcService) setEnv(key string, value string)  {
	for k, v := range s.Environment {
		if v.Key == key {
			s.Environment[k].Value = value
			return
		}
	}
	s.addEnv(key,value)
}

type Logging struct{
	Driver string `json:"driver,omitempty"`
}

type UnixEnvString struct {
	Key string
	Value string
}

func (ues *UnixEnvString) UnmarshalJSON(data []byte) error {
	raw := string(data)
	split := strings.Split(raw, "=")
	if len(split) < 2 {
		return errors.New("bad UnixEnvString")
	}
	ues.Key = strings.Trim(split[0],"\"")
	ues.Value = strings.Trim(split[1],"\"")
	return nil
}

func (ues *UnixEnvString) MarshalJSON() ([]byte, error) {
	return []byte( "\"" + ues.Key + "=" + ues.Value + "\""), nil
}