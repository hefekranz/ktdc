package ktdc

import (
	"encoding/json"
	"github.com/ghodss/yaml"
	"errors"
)

type Unmarshal func(data []byte, ref interface{}) error

type ReadFile func(path string) ([]byte, error)

type FileLoader struct {
	readFile ReadFile
}

func NewFileLoader(reader ReadFile) *FileLoader {
	fileLoader := new(FileLoader)
	fileLoader.readFile = reader
	return fileLoader
}

func (fl FileLoader) LoadConfig(path string) (Config, error) {
	config := Config{}
	err := fl.CreateFromFile(path, &config, json.Unmarshal)
	return config, err
}

func (fl FileLoader) LoadDcFile(path string) (DcFile, error) {
	dcFile := DcFile{}
	err := fl.CreateFromFile(path, &dcFile, yaml.Unmarshal)
	return dcFile, err
}

func (fl FileLoader) LoadDeploymentFile(path string) (KubeDeployment,error) {
	kubeDeployment := KubeDeployment{}
	err := fl.CreateFromFile(path,&kubeDeployment,yaml.Unmarshal)
	return kubeDeployment, err
}

func (fl FileLoader) LoadDeploymentMapFromConfig(config Config) (map[string]KubeDeployment, error)  {
	deploymentMap := make(map[string]KubeDeployment)
	for name, serviceConfig := range config.Services {
		if len(serviceConfig.DeploymentFile) == 0 {
			return deploymentMap, errors.New("missing DeploymentFile for service " + name)
		}
		deployment, err := fl.LoadDeploymentFile(serviceConfig.DeploymentFile)
		if err != nil { return deploymentMap, err}

		deploymentMap[name] = deployment
	}
	return deploymentMap, nil
}

func (fl FileLoader) CreateFromFile(path string, reference interface{}, unmarshal Unmarshal) error {
	raw, err := fl.readFile(path)
	if err != nil {
		return err
	}

	return unmarshal(raw, &reference)
}
