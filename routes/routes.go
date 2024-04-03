package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jpcadinelli/curso-go-gin-alura-api/controllers"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	r.GET("/:nome", controllers.Saudacao)
	r.Run()
}