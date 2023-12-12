package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Piszmog/cloudconfigclient/v2"
)

func main() {
	configClient, err := cloudconfigclient.New(cloudconfigclient.Local(&http.Client{}, "http://localhost:8888"))

	if err != nil {
		fmt.Println(err)
		return
	}
	// var file File
	// // Retrieves a 'temp1.json' from the Config Server's default branch in directory 'temp' and deserialize to File
	// err = configClient.GetFile("temp", "temp1.json", &file)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Printf("%+v\n", file)
	// Retrieves a 'temp2.txt' from the Config Server's default branch in directory 'temp' as a byte slice ([]byte)
	b, err := configClient.GetFileFromBranchRaw("master", "playservice/dev", "flags.yaml")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))

	lastVersion := ""
	for i := 0; i < 6; i++ {
		// Retrieves the configurations from the Config Server based on the application name and active profiles
		fmt.Println("Getting Config")
		config, err := configClient.GetConfiguration("playservice", "dev/master")
		if err != nil {
			fmt.Println(err)
			return
		}
		if config.Version != lastVersion {
			fmt.Printf("New Version %+v\n", config)
			// if we want, we can convert the config to a struct
			var configStruct Config
			err = config.Unmarshal(&configStruct)
			if err != nil {
				fmt.Println(err)
			}
		}
		lastVersion = config.Version
		time.Sleep(10 * time.Second)
	}
}

type Config struct {
	ServiceConfig  string `json:"service-config"`
	Flags          string `json:"flags"`
	RuleParameters string `json:"rule-params"`
}
