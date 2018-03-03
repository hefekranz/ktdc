package ktdc

import "log"

type Converter interface {
	convert()
}

func Convert(kubeFiles map[string]KubeDeployment, dcFile *DcFile, config Config) {

	for name := range config.Services {

		deployment, ok := kubeFiles[name]
		if !ok {
			log.Println("key " + name + " not found")
			continue
		}

		service := buildServiceFromKubernetes(deployment)

		dcFile.addService(name, addConfig(name,service,config))
	}
}

func buildServiceFromKubernetes(deployment KubeDeployment) DcService{
	service := DcService{}
	service.Image = deployment.getImage()
	for _, v := range deployment.getEnvs() {
		service.addEnv(v.Name, v.Value)
	}
	return service
}

func addConfig(name string ,service DcService,config Config)  DcService{

	for k,v := range config.GlobalEnvs {
		service.overwriteEnv(k,v)
	}

	serviceConfig, ok := config.Services[name]
	if ok {
		for k,v := range serviceConfig.Envs {
			service.setEnv(k,v)
		}

		service.Ports = serviceConfig.Ports
		service.Command = serviceConfig.Command
	}

	return service
}
