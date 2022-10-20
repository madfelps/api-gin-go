package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/madfelps/api-gin-go/controllers"
	"github.com/madfelps/api-gin-go/database"
	"github.com/madfelps/api-gin-go/models"
	"github.com/stretchr/testify/assert"
)

func SetupDasRotas() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	rotas := gin.Default()
	return rotas
}

func TestVerificaStatusCodeDaSaudacaoComParametro(t *testing.T) {
	r := SetupDasRotas()
	r.GET("/:nome", controllers.Saudacao)
	req, _ := http.NewRequest("GET", "/fe", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	/* if resposta.Code != http.StatusOK {
		t.Fatalf("Status error: valor recebido foi %d e o esperado era %d", resposta.Code, http.StatusOK)
	} */
	assert.Equal(t, http.StatusOK, resposta.Code, "Deveriam ser iguais")
	mockDaResposta := `{"API diz:":"E aí fe, tudo beleza?"}`
	respostaBody, _ := ioutil.ReadAll(resposta.Body)
	assert.Equal(t, mockDaResposta, string(respostaBody))

}

var ID int

func CriaAlunoMock() {
	aluno := models.Aluno{Nome: "Nome do Aluno Teste", CPF: "12345678901", RG: "123456789"}
	database.DB.Create(&aluno)
	//ID := int(aluno.ID)

}

func DeletaAlunoMock() {
	var aluno models.Aluno
	database.DB.Delete(&aluno, ID)
}

func TestListandoTodosOsAlunosHandler(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupDasRotas()
	r.GET("/alunos", controllers.ExibeTodosOsAlunos)
	req, _ := http.NewRequest("GET", "/alunos", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
	//fmt.Println(resposta.Body)
}

func TestBuscaAlunoPorCPFHandler(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupDasRotas()
	r.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoPorCPF)
	req, _ := http.NewRequest("GET", "/alunos/cpf/12345678901", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
}
