package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/cjcs19/demotocidenet/bd"
	"github.com/cjcs19/demotocidenet/utils"
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
    // curl -i -X POST -H "Content-Type: application/json" -d "{ \"NombrePais\": \"EEUU\"}" http://localhost:10000/api/v2/pais
}

func PostIdentificacion(c *gin.Context) {
    db := bd.InitDb()
      defer db.Close()
   var identificacion bd.TipoIdentificacion
   c.Bind(&identificacion)
   
   if identificacion.NombreTipo != ""  {
       // INSERT INTO "users" (name) VALUES (user.Name);
       db.Create(&identificacion)
       // Display error
       c.JSON(201, gin.H{"success": identificacion})
   } else {
       // Display error
       c.JSON(422, gin.H{"error": "Fields are empty"})
   }
    // curl -i -X POST -H "Content-Type: application/json" -d "{ \"NombreTipo\": \"Pasaporte\"}" http://localhost:10000/api/v2/identificacion
}

func GetPaises(c *gin.Context){
	db := bd.InitDb()
      defer db.Close()
   var pais []bd.PaisEmp 

   db.Find(&pais)

   c.JSON(200, pais)

   // curl -i http://localhost:10000/api/v2/paises

}

func PostUser(c *gin.Context){
	db := bd.InitDb()
	defer db.Close()
 	var users bd.Users

	c.Bind(&users)

	//creacion de correo
	users.Correo=utils.Correo(users)



}