package main

import (
	"github.com/cjcs19/demotocidenet/controllers"
	"github.com/gin-gonic/gin"
	//"github.com/jinzhu/gorm"
	//_ "github.com/jinzhu/gorm/dialects/postgres"
)




func main() {
    r := gin.Default()

    //v1 := r.Group("api/v1")
    //{
    //    v1.POST("/users", PostUser)
    //    v1.GET("/users", GetUsers)
    //    v1.GET("/users/:id", GetUser)
    //    v1.PUT("/users/:id", UpdateUser)
    //    v1.DELETE("/users/:id", DeleteUser)
    //}

    v2 := r.Group("api/v2")
    {
        v2.POST("/pais", controllers.PostPais)
        v2.POST("/identificacion", controllers.PostIdentificacion)
        v2.GET("/pais", controllers.GetPaises)        
    }

    r.Run(":10000")
}