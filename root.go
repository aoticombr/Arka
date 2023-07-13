package main

import (
	"encoding/base64"
	"fmt"

	"github.com/luraproject/lura/v2/config"
)

var (
	DefaultRoot Root
	cfgFile     string
	debug       int
	port        int
	parser      config.Parser
	run         func(config.ServiceConfig)
)

func init() {
	logo, err := base64.StdEncoding.DecodeString(encodedLogo)
	if err != nil {
		fmt.Println("decode error:", err)
	}
	fmt.Println(string(logo))

}

const encodedLogo = "CiDilojilojilojilojilojilZcg4paI4paI4paI4paI4paI4paI4pWXIOKWiOKWiOKVlyAg4paI4paI4pWXIOKWiOKWiOKWiOKWiOKWiOKVlyAK4paI4paI4pWU4pWQ4pWQ4paI4paI4pWX4paI4paI4pWU4pWQ4pWQ4paI4paI4pWX4paI4paI4pWRIOKWiOKWiOKVlOKVneKWiOKWiOKVlOKVkOKVkOKWiOKWiOKVlwrilojilojilojilojilojilojilojilZHilojilojilojilojilojilojilZTilZ3ilojilojilojilojilojilZTilZ0g4paI4paI4paI4paI4paI4paI4paI4pWRCuKWiOKWiOKVlOKVkOKVkOKWiOKWiOKVkeKWiOKWiOKVlOKVkOKVkOKWiOKWiOKVl+KWiOKWiOKVlOKVkOKWiOKWiOKVlyDilojilojilZTilZDilZDilojilojilZEK4paI4paI4pWRICDilojilojilZHilojilojilZEgIOKWiOKWiOKVkeKWiOKWiOKVkSAg4paI4paI4pWX4paI4paI4pWRICDilojilojilZEK4pWa4pWQ4pWdICDilZrilZDilZ3ilZrilZDilZ0gIOKVmuKVkOKVneKVmuKVkOKVnSAg4pWa4pWQ4pWd4pWa4pWQ4pWdICDilZrilZDilZ0KICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAK"
