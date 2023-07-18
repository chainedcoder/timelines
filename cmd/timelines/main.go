package main

import (
	"chainedcoder/timelines/cmd/config"
	"chainedcoder/timelines/internal/logger"

	"chainedcoder/timelines/internal/orm"
	"chainedcoder/timelines/pkg/server"
)

// main
func main() {
	sc := config.Server()
	orm, err := orm.Factory(sc)
	defer orm.DB.Close()
	if err != nil {
		logger.Panic(err)
	}
	server.Run(sc, orm)
}