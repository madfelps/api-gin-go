package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/madfelps/api-gin-go/controllers"
	controller "github.com/madfelps/api-gin-go/controllers"
)

func HandleRequests() {
	r := gin.Default()

	r.GET("/alunos", controller.ExibeTodosOsAlunos)
	r.GET("/:nome", controller.Saudacao)
	r.POST("/alunos", controllers.CriaNovoAluno)
	r.GET("/alunos/:id", controllers.BuscaAlunoPorID)
	r.DELETE("/alunos/:id", controllers.DeletaAluno)
	r.PATCH("/alunos/:id", controllers.EditaAluno)
	r.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoPorCPF)
	r.Run(":5300")
}
