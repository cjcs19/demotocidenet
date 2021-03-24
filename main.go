package main

import (
	"github.com/cjcs19/demotocidenet/controllers"
	"github.com/gin-gonic/gin"
	
)




func main() {
    r := gin.Default()

    v1 := r.Group("api/v1")
    {
        v1.POST("/users", controllers.PostUser)
    }    

    v2 := r.Group("api/v2")
    {
        v2.POST("/pais", controllers.PostPais)
        v2.POST("/identificacion", controllers.PostIdentificacion)
        v2.GET("/pais", controllers.GetPaises)        
    }

    r.Run(":10000")
}