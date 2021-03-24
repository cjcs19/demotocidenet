package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/cjcs19/demotocidenet/bd"
)

func PostPais(c *gin.Context) {
    db := bd.InitDb()
      defer db.Close()
   var pais bd.PaisEmp
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

