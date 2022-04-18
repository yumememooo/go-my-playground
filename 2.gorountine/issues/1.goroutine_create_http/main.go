package main

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

//這個範例用來試著創建大量的go，然後透過gin的pprof接口監看效能
func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.POST("/login", loginEndpoint)
	}
	engine := gin.New()
	pprof.Register(engine) // pprof monitor
	endpoint := fmt.Sprintf(":%d", 8888)
	go func() { //分析用
		err := engine.Run(endpoint)
		if err != nil {
			fmt.Println(err)
		}
	}()

	fmt.Printf("end call")
	select {} //避免退出
}
func loginEndpoint(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
}
