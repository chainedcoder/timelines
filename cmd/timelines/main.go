package main

import (
	log "chainedcoder/timelines/internal/logger"

	"chainedcoder/timelines/internal/orm"
	"chainedcoder/timelines/pkg/server"
)

func main() {
  // Create a new ORM instance to send it to our
  orm, err := orm.Factory()
  if err != nil {
    log.Panic(err)
  }
  // Send: ORM instance
  server.Run(orm)
}