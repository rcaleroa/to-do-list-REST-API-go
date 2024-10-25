package main

import(
	"net/http"
	"github.com/gin-gonic/gin"
)

type PorHacer struct{
	Hecho bool
	Nombre string
	Descripcion string
}


func main(){
	router := gin.Default()

	router.GET("/porhacer", func(c *gin.Context){
		c.IndentedJSON(http.StatusOK, []PorHacer{
			//lista de tareas por hacer
			{
				Nombre : "shower",
			},
			{
				Nombre : "cook & eat breakfast",
			},
			{
				Nombre : "code",
			},
			{
				Nombre : "play ps5",
			},
			{
				Nombre : "code",
			},
			{
				Nombre : "make & eat lunch",
			},
			{
				Nombre : "code",
			},
			{
				Nombre : "sleep",
			},
		})

	})

	router.Run("localhost:8080")
}