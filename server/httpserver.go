package server

import (
	v1 "gomedia/server/v1"
	v2 "gomedia/server/v2"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/semaphore"
)

const concurrency = 500

var sem *semaphore.Weighted

func RunServer() {
	r := gin.Default()

	//Setup the router
	setupRouter(r)

	//Config the server and cors
	server := ServerCustom(r)
	corsConfig := UpdateCors()
	r.Use(cors.New(corsConfig))

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

func setupRouter(r *gin.Engine) {
	sem = semaphore.NewWeighted(concurrency)

	r.Use(ConcurrencyLimitMiddleware)

	v1.RegisterUserRoutesV1(r.Group("/v1"))

	v2.RegisterUserRoutesV2(r.Group("/v2"))

	//For SPA frontend to work like React, Angular, Vue
	r.Static("/static", "./public")

	//Serve the video hls files
	r.StaticFS("/videos", http.Dir("/videos"))

	//Handle not found routes
	r.NoRoute(handleNotFound)
}

func handleNotFound(c *gin.Context) {
	c.JSON(404, gin.H{"message": "Not found"})
}
