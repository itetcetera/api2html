package main

//go:generate statik -src=skeleton/files

import (
	"log"

	"github.com/devopsfaith/api2html/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Println("error:", err.Error())
	}
}
