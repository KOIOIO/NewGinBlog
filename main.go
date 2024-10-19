package main

import (
	"NewGinBlog/Model"
	"NewGinBlog/Routers"
	"NewGinBlog/Utills"
	"log"
)

func main() {
	log.Printf("Listening and serving HTTP on %s\n", Utills.HttpPort)
	Model.InitDb()
	Routers.InitRouter()

}
