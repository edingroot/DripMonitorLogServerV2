// Ref: http://goinbigdata.com/persisting-application-configuration-in-golang/
package helpers

import (
	"os"
	"encoding/json"
	"io/ioutil"
	"fmt"
)

// Path is related to project root
const relativeConfigPath = "/config/config.json"
const runtimeEnvKey = "RUNTIME_ENV"

// Instances should be created by calling NewConfiguration()
type Configuration struct {
	runtimeEnv
}

type configRoot struct {
	Production   runtimeEnv `json:"production"`
	Development  runtimeEnv `json:"development"`
	Test  		 runtimeEnv `json:"test"`
}

type runtimeEnv struct {
	HttpListenHost    string  `json:"http_listen_host"`
	HttpListenPort    int     `json:"http_listen_port"`
	StreamListenHost  string  `json:"stream_listen_host"`
	StreamListenPort  int     `json:"stream_listen_port"`
	DbHost 	          string  `json:"db_host"`
	DbPort            int     `json:"db_port"`
	DbUsername        string  `json:"db_username"`
	DbPassword        string  `json:"db_password"`
	DbName            string  `json:"db_name"`
}

func NewConfiguration() (*Configuration, error) {
	bytes, err := ioutil.ReadFile(getAbsConfigPath())
	if err != nil {
		return nil, err
	}

	var config Configuration
	var configRoot configRoot
	err = json.Unmarshal(bytes, &configRoot)
	if err != nil {
		return nil, err
	}

	var mode string = os.Getenv(runtimeEnvKey)
	switch mode {
	case "production":
		config.runtimeEnv = configRoot.Production

	case "development":
		config.runtimeEnv = configRoot.Development

	case "test":
		config.runtimeEnv = configRoot.Test

	default:
		// TODO: use logging library
		fmt.Println("Warning: " + runtimeEnvKey + " not set.")
		mode = "production"
		config.runtimeEnv = configRoot.Production
	}
	fmt.Println("Running on " + mode + " mode.")

	return &config, nil
}


/*
func SaveConfig(config *Configuration) error {
	bytes, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(getAbsConfigPath(), bytes, 0644)
}
*/

func getAbsConfigPath() string {
	return ProjectPath + relativeConfigPath
}
