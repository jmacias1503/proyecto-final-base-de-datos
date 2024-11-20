package main

import (
  "os"
  "net/http"
  "gorm.io/gorm"
  "gorm.io/driver/postgres"
  "github.com/gin-gonic/gin"
)

type Personas struct {
  gorm.Model
  Curp string
  Nombre string
  Primerapellido string
  Segundoapellido string
  Nacionalidad string
  Sexo bool
  Lugarnacimiento string
  Fechanacimiento string
  Estadocivil string
}

func main(){
  dsn := os.Getenv("DATABASE_URL")
  db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
    SkipDefaultTransaction: true,
  })
  if err != nil {
    panic("Error connecting to db")
  }
  db.AutoMigrate(&Personas{})

  router := gin.Default()
  router.LoadHTMLGlob("templates/*")

  router.GET("/", func(ctx *gin.Context) {
    ctx.HTML(http.StatusOK, "index.html", gin.H{
    })
  })
  router.POST("/", func(ctx *gin.Context) {
    var person Personas
    if err := ctx.ShouldBind(&person); err == nil {
      db.Create(&person)
      ctx.JSON(http.StatusCreated, person)
    } else {
      ctx.JSON(http.StatusBadRequest, gin.H{
	"error": "Invalid payload",
      })
    }
  })
  router.Run(":"+os.Getenv("WORKING_PORT"))
}
