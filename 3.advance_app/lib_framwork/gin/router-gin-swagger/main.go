package main

import "go-demo/3.advance_app/lib_framwork/gin/router-gin-swagger/router"

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
func main() {
	r := router.SetupRouter()
	r.Run() // default localhost:8000
}
