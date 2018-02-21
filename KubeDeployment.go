package ktdc

type KubeDeployment struct {
	Spec struct {
		Template struct{
			Spec struct{
				Containers []struct{
					Image string
					Env []KubeEnvEntry
				}
			}
		}
	}
}

func (kd KubeDeployment) getImage() string  {
	return kd.Spec.Template.Spec.Containers[0].Image
}

func (kd KubeDeployment) getEnvs() []KubeEnvEntry {
	return kd.Spec.Template.Spec.Containers[0].Env
}

type KubeEnvEntry struct {
	Name string
	Value string
}