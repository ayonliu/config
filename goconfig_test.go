package goconfig

import (
	"fmt"
	"testing"
)

type Site struct {
	Title string `mapstructure:"title"`
}

type DB struct {
	Host     string `mapstructure:"host"`
	Name     string `mapstructure:"name"`
	Password string `mapstructure:"password"`
	Port     string `mapstructure:"port"`
	Username string `mapstructure:"username"`
}

type MyConfig struct {
	Db   DB
	Site Site
}

func TestGetConfig(t *testing.T) {
	var conf MyConfig
	runtimeConf, err := GetConfig(conf)
	if err != nil {
		t.Fatalf("cannot load config:%v", err)
	}
	fmt.Printf("runtimeConf: %T\n", runtimeConf)
	fmt.Printf("%+v\n", runtimeConf)
	// see https://go.dev/doc/effective_go#interface_conversions
	confStruct, ok := runtimeConf.(MyConfig)
	if !ok {
		t.Fatalf("runtimeConf cannot convert to MyConfig")
	}
	fmt.Printf("confStruct: %+v\n", confStruct)
	fmt.Printf("Db.Username:%v\n", confStruct.Db.Username)
}
