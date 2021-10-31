package config

import (
	"log"

	"github.com/zerodays/gonfig"
	"github.com/zerodays/gonfig/providers/environment"
	"github.com/zerodays/gonfig/providers/ini"
)

var cfg *gonfig.Config

func Load() {
	data, err := Asset("config.ini")
	if err != nil {
		log.Fatal(err)
	}

	iniProvider, err := ini.New(data)
	if err != nil {
		log.Fatal(err)
	}

	iniProvider.SectionMode = ini.ModeLowercase
	iniProvider.KeyMode = ini.ModeUppercase

	cfg = gonfig.New(iniProvider, environment.Provider{})
	cfg.AppName = "prpo-auth"

	err = Login.LoadKeys()
	if err != nil {
		log.Printf("Could not load signing keys (error: %s). You can generate them with `genkeys` command.\n",
			err.Error())
	}
}
