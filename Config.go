package ktdc

type Config struct {
	GlobalEnvs map[string]string
	Services map[string]ServiceConfig
}

type ServiceConfig struct {
	Repository string
	DeploymentFile string `json:"deploymentFile"`
	Envs map[string]string
}