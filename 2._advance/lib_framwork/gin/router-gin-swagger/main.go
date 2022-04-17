package main

import "go-demo/2._advance/lib_framwork/gin/router-gin-swagger/router"

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
func main() {
	r := router.SetupRouter()
	r.Run() // default localhost:8000
}
