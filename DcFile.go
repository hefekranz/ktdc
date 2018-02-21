package ktdc

type DcFile struct {
	Version string                `json:"version"`
	Services map[string]DcService `json:"services"`
}

func NewDcFile() DcFile {
	dcFile := DcFile{}
	dcFile.Version = "2"
	dcFile.Services = make(map[string]DcService)
	return dcFile
}

func (f DcFile) addService(name string, service DcService)  {
	f.Services[name] = service
}
