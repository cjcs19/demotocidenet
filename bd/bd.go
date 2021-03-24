package bd

import (     
	  _ "github.com/jinzhu/gorm/dialects/postgres"
	  "github.com/jinzhu/gorm"
)

type PaisEmp struct {
    Id          int `gorm:"AUTO_INCREMENT" form:"id" json:"id"`
    NombrePais  string
}

type TipoIdentificacion struct {
    Id                      int `gorm:"AUTO_INCREMENT" form:"id" json:"id"`
    NombreTipo              string
}
type Users struct {
    Id                       int                   `gorm:"AUTO_INCREMENT" form:"id" json:"id"`    
    Nombre                   string                `gorm:"not null" form:"nombre" json:"nombre"`
    OtrosNombres             string                `gorm:"not null" form:"otrosnombres" json:"otrosnombres"`
    PrimerApellido           string                `gorm:"not null" form:"primerapellido" json:"primerapellido"`
    SegundoApellido          string                `gorm:"not null" form:"segundoapellido" json:"segundoapellido"`
    Idpais                   int                   `gorm:"not null" form:"idpais" json:"idpais"`
    IdTipoIdent              int                   `gorm:"not null" form:"idtipoident" json:"idtipoident"`
    Correo                   string                `gorm:"not null" form:"correo" json:"correo"`

}

func InitDb() *gorm.DB {
    
	db, err := gorm.Open("postgres", "host=127.0.0.1 user=cidenet dbname=cidenetdb sslmode=disable password=cidenet2020")
  
    db.LogMode(true)
	
    // Error
    if err != nil {
        panic(err)
    }

    // Creating the table
    db.AutoMigrate(&Users{})
    db.AutoMigrate(&PaisEmp{})
    db.AutoMigrate(&TipoIdentificacion{})

    return db
}