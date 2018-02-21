package ktdc

import "log"

type Converter interface {
	convert()
}

func Convert(kubeFiles map[string]KubeDeployment, dcFile *DcFile, config Config) {

	for k := range config.Services {

		deployment, ok := kubeFiles[k]
		if !ok {
			log.Println("key " + k + " not found")
			continue
		}

		service := buildServiceFromKubernetes(k,deployment)

		dcFile.addService(k, addConfig(service,config))
	}
}

func buildServiceFromKubernetes(name string, deployment KubeDeployment) DcService{
	service := DcService{}
	service.Image = deployment.getImage()
	service.ContainerName = name
	for _, v := range deployment.getEnvs() {
		service.addEnv(v.Name, v.Value)
	}
	return service
}

func addConfig(service DcService,config Config)  DcService{

	for k,v := range config.GlobalEnvs {
		service.overwriteEnv(k,v)
	}

	serviceConfig, ok := config.Services[service.ContainerName]
	if ok {
		for k,v := range serviceConfig.Envs {
			service.setEnv(k,v)
		}

		service.Ports = serviceConfig.Ports
	}

	return service
}
