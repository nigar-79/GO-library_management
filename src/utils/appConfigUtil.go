package utils

import (
	"fmt"
	"os"

	"github.com/pelletier/go-toml"
	"gorm.io/gorm"
	"sdbgo.in/sdbgo/src/config"
)

//DatabaseConn is the struct used to initilaise the Appconfig struct from the config file which holds the configuration structs.
type DatabaseConn struct {
	AppConfiguration *config.Database
	Data             *gorm.DB
}

//LoadAppConfig Function loads the toml file data to the struct defined in appconfig and
//returns an error with the data
func LoadAppConfig(filepath string) (*config.AppConfig, error) {

	var appConfig config.AppConfig

	if filepath == "app.dev.toml" {
		file, err := os.Open(filepath)
		if err != nil {
			return &appConfig, fmt.Errorf(err.Error())
		}
		defer file.Close()
		tomlParser := toml.NewDecoder(file)
		err = tomlParser.Decode(&appConfig)
		if err != nil {
			return nil, fmt.Errorf(err.Error())
		}
	} else if filepath == "app.test.toml" {
		file, err := os.Open(filepath)
		if err != nil {
			return &appConfig, fmt.Errorf(err.Error())
		}
		tomlParser := toml.NewDecoder(file)
		err = tomlParser.Decode(&appConfig)
		if err != nil {
			return nil, fmt.Errorf(err.Error())
		}
	} else if filepath == "app.prod.toml" {
		file, err := os.Open(filepath)
		if err != nil {
			return &appConfig, fmt.Errorf(err.Error())
		}
		defer file.Close()
		tomlParser := toml.NewDecoder(file)
		err = tomlParser.Decode(&appConfig)
		if err != nil {
			return nil, fmt.Errorf(err.Error())
		}
	} else {
		fmt.Println("File not found")
	}

	return &appConfig, nil
}
