package router

import (

	/* "./docs" // docs is generated by Swag CLI, you have to import it.
	docs "github.com/go-project-name/docs" 需要對應到你自己的project
	docs "go-demo/2._advance/lib_framwork/gin/router/docs" //改成這樣 go.mod必須是一樣的module name,不可以用main
	*/
	docs "go-demo/3.advance_app/lib_framwork/gin/router-gin-swagger/docs"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// 回傳的 Type 是 *gin.Engine
func SetupRouter() *gin.Engine {
	// Disable log's color
	gin.DisableConsoleColor()

	r := gin.Default()
	r.NoMethod(HandleNotFound)
	r.NoRoute(HandleNotFound)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler)) //http://localhost:8080/swagger/index.html

	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	r.GET("/getJson", getJson)

	//Group example
	r.GET("/users/:name", userByname) //path Parameters
	v1 := r.Group("/v1")
	{
		v1.POST("/login", loginEndpoint)
	}
	return r

}

func userByname(c *gin.Context) {
	name := c.Param("name")
	c.String(http.StatusOK, "Hello %s", name)
}

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /getJson [get]
func getJson(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":  "posted",
		"message": "message",
		"nick":    "nick",
	})
}
func loginEndpoint(c *gin.Context) {
	var json Login
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if json.User != "manu" || json.Password != "123" {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
}

type Login struct {
	User     string `form:"user" json:"user" xml:"user"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func HandleNotFound(c *gin.Context) {
	handleErr := NotFound()
	handleErr.Request = c.Request.Method + " " + c.Request.URL.String()
	c.JSON(handleErr.Code, handleErr)
	return
}

