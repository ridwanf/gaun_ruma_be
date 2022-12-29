package main

import (
	"gaunRumaRestApi/config"
	"gaunRumaRestApi/config/db"
	"gaunRumaRestApi/routes"
	"log"
)

func main() {
	conf, err := config.GetConfig(".")
	if err != nil {
		log.Fatal("cannot load config ", err)
	}
	d := db.Init(conf)
	e := routes.Init(d, conf)

	e.Logger.Info(e.Start(":3000"))
}
