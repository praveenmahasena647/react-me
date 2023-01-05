package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/praveenmahasena647/todo/routes"
)

func main() {
	var router *gin.Engine = routes.RunServer()

	var err error = router.Run("localhost:42069")

	if err != nil {
		log.Println(err.Error())
		os.Exit(2)
	}
}
