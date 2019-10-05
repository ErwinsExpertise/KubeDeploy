package client

import (
	"log"

	"github.com/BurntSushi/toml"
)

type Database struct {
	Host        string
	Username    string
	Password    string
	StorageName *string
}

func (d *Database) BuildConfig() {
	if _, err := toml.DecodeFile("config.toml", &d); err != nil {
		log.Printf("\nError building config: \n%+v", err)
	}

}

