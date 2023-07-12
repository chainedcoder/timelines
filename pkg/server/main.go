package server

import (
	"log"

	"chainedcoder/timelines/internal/handlers"
	"chainedcoder/timelines/pkg/utils"

	"github.com/gin-gonic/gin"
)

var host, port string

func init() {
    host = utils.MustGet("TIMELINES_SERVER_HOST")
    port = utils.MustGet("TIMELINES_SERVER_PORT")
}

// Run web server
func Run() {
    r := gin.Default()
    // Setup routes
    r.GET("/ping", handlers.Ping())
    log.Println("Running @ http://" + host + ":" + port )
    log.Fatalln(r.Run(host + ":" + port))
}