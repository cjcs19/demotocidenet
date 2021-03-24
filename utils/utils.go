package utils

import (
	
	"github.com/cjcs19/demotocidenet/bd"
	"fmt"
	
)

func Correo(usr users) string  {
	
	return usr.Nombre + '.' + usr.PrimerApellido + '@' + 'cidenet.com'

	
}