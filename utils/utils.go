package utils

import (
	
	"github.com/cjcs19/demotocidenet/bd"
	"strconv"
	//"fmt"
	
)

func Correo(usr *bd.Users)   {
	db := bd.InitDb()
	defer db.Close()

	var user bd.Users

	result := db.Where("nombre = ? AND primer_apellido = ? ", usr.Nombre,usr.PrimerApellido).Find(&user) 

	cantidad:=strconv.Itoa(int(result.RowsAffected +1))


	dominio:="cidenet.com.co"
	if usr.Idpais==1 {
		dominio="cidenet.com.us"
	}

	usr.Correo=usr.Nombre + "." + usr.PrimerApellido + "." + cantidad + "@" + dominio 
	

}