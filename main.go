package main

import (
	"github.com/cjcs19/demotocidenet/bd"
	"github.com/gin-gonic/gin"
	//"github.com/jinzhu/gorm"
	//_ "github.com/jinzhu/gorm/dialects/postgres"
)

func PostPais(c *gin.Context) {
    db := bd.InitDb()
      defer db.Close()
   var pais PaisEmp
   c.Bind(&pais)
   
   if pais.NombrePais != ""  {
       // INSERT INTO "users" (name) VALUES (user.Name);
       db.Create(&pais)
       // Display error
       c.JSON(201, gin.H{"success": pais})
   } else {
       // Display error
       c.JSON(422, gin.H{"error": "Fields are empty"})
   }
    // curl -i -X POST -H "Content-Type: application/json" -d "{ \"firstname\": \"Thea\", \"lastname\": \"Queen\" }" http://localhost:8080/api/v1/users
}


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
        v2.POST("/pais", PostPais)
        //v2.GET("/pais", GetPais)        
    }

    r.Run(":10000")
}