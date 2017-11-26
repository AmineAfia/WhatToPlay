package config

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

type ConfigValues struct {
	DatabaseConnectString  string
	BaseUrl                string
	AppUrl                 string
	SpotifyAppClientId     string
	SpotifyAppClientSecret string
	SpotifyClient          string
}

var Conf = DefaultConfig()

func DefaultConfig() ConfigValues {
	fmt.Print("Initializing configs\n")

	values := ConfigValues{}
	values.DatabaseConnectString = "mongodb://admin:<PASSWORD>@localhost"
	values.BaseUrl = "http://localhost:8080/"
	values.AppUrl = "http://localhost:3000/"
	values.SpotifyAppClientId = "get this shit from spotify dev portal"
	values.SpotifyAppClientSecret = "this shit too"
	values.SpotifyClient = "yourUserId"

	return values
}

func LoadConfiguration(cfgFile string) {
	var conf ConfigValues

	tomlData, err := ioutil.ReadFile(cfgFile)

	log.Println("Loading Config")

	if string(tomlData) == "" {
		SaveConfiguration(cfgFile)
		return
	}

	if err != nil {
		log.Println(err)
		Conf = DefaultConfig()
		SaveConfiguration(cfgFile)
		return
	}

	if _, err := toml.Decode(string(tomlData), &conf); err != nil {
		log.Println(err)
		SaveConfiguration(cfgFile)
	} else {
		Conf = conf
	}

}

func SaveConfiguration(cfgFile string) {

	log.Println("Saving Config")

	file, err := os.Create(cfgFile)
	if err != nil {
		log.Println(err)
		panic(err)
	}

	defer file.Close()

	w := bufio.NewWriter(file)
	enc := toml.NewEncoder(w)
	err = enc.Encode(Conf)

	if err != nil {
		log.Println(err)
		panic(err)
	}
}
