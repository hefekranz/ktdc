package main

import (
	"log"
	"io/ioutil"
	"fmt"
	"github.com/ghodss/yaml"
	"flag"
	"hefekranz/ktdc"
	"os"
)

func handleError(e error) error{
	if e != nil {
		log.Fatal(e)
	}
	return e
}

func init()  {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main()  {

	fileLoader := ktdc.NewFileLoader(ioutil.ReadFile)

	configPath := flag.String("config","","path to config.json")
	dcFilePath := flag.String("dcFile","", "docker-compose file to add to")
	outputFile := flag.String("output","docker-compose.yaml","output file name")
	flag.Parse()

	if empty(string(*configPath)){
		flag.Usage()
		log.Fatal("missing config.json")
	}

	config, err := fileLoader.LoadConfig(*configPath)
	handleError(err)

	var dcFile = ktdc.NewDcFile()
	if !empty(*dcFilePath) {
		dcFile, err = fileLoader.LoadDcFile(*dcFilePath)
		handleError(err)
	}

	kubeDeployments, err := fileLoader.LoadDeploymentMapFromConfig(config)
	handleError(err)

	ktdc.Convert(kubeDeployments,&dcFile,config)

	output, err := yaml.Marshal(&dcFile)
	handleError(err)
	fmt.Println("done!")
	ioutil.WriteFile(*outputFile,output, os.ModePerm)
}

func empty(s string) bool {
	return len(s) == 0
}