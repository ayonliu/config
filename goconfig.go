package goconfig

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

func GetConfig(confStruct interface{}, config ...string) (interface{}, error) {
	var (
		config_name = "config"
		config_type = "yml"
	)
	if len(config) > 0 {
		fmt.Println(config[0])
		if strings.Contains(config[0], ".") {
			configs := strings.Split(config[0], ".")
			config_name = configs[0]
			config_type = configs[1]
		}
	}
	// Set the file name of the configurations file
	viper.SetConfigName(config_name)

	// Set the path to look for the configurations file
	viper.AddConfigPath(".")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	viper.SetConfigType(config_type)

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	/*
		// Set undefined variables
		viper.SetDefault("DB.HOST", "127.0.0.1")

		// getting env variables DB.PORT
		// viper.Get() returns an empty interface{}
		// so we have to do the type assertion, to get the value
		DBPort, ok := viper.Get("DB.PORT").(string)

		// if type assert is not valid it will throw an error
		if !ok {
			log.Fatalf("Invalid type assertion")
		}

		fmt.Printf("viper : %s = %s \n", "Database Port", DBPort)
	*/

	// unmarshal config
	err := viper.Unmarshal(&confStruct)

	return confStruct, err
}
