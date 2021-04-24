package main

import "carrot-system/core/src/config"
import "github.com/gin-gonic/gin"

func main() {
	config.LoadEnv()

	r := gin.Default()

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}