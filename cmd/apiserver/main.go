package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"log"
	"regappeals/internal/app/apiserver"
)

var (
	configPath string //nolint:gochecknoglobals
)

func main() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
	flag.Parse()

	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	if err = apiserver.Start(config); err != nil {
		log.Println(err)
	}
}
