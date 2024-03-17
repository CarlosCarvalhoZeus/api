package handlers

import (
	"api/db"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetPessoas(c *gin.Context) {

	con, err := db.Conn()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}
	defer con.Close()
	pessoas, err := db.GetPessoas(con)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}
	c.JSON(http.StatusOK, pessoas)
}

func GetProdutos(c *gin.Context) {
	con, err := db.Conn()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}
	defer con.Close()
	produtos, err := db.GetProdutos(con)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}
	c.JSON(http.StatusOK, produtos)
}

func GetInsightPessoa(c *gin.Context) {
	con, err := db.Conn()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}
	defer con.Close()
	id := c.Param("id")
	pes, err := db.GetPessoasInsigh(con, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}
	// var msg []byte
	var pesRes []db.ProdutoFeedback
	var sumOkBom int64
	var sumRuim int64

	var recomendacao []string
	for _, v := range pes {

		var aux db.ProdutoFeedback
		aux.Bom = v.Bom
		aux.Ruim = v.Ruim
		aux.Indiferente = v.Indiferente

		sumOkBom += v.Indiferente + v.Bom
		sumRuim += v.Ruim
		pesRes = append(pesRes, aux)
		if v.Indiferente+v.Bom > v.Ruim {
			if v.Bom > v.Indiferente {
				recomendacao = append(recomendacao, fmt.Sprintf("Recomendado para tarefas envolvendo produto: %v", v.Produto))
				continue
			}
			recomendacao = append(recomendacao, fmt.Sprintf("Pode realizar tarefas envolvendo produto: %v, porem não é o mais recomendado", v.Produto))
			continue
		}
		recomendacao = append(recomendacao, fmt.Sprintf("Não é recomendado para realizar tarefas envolvendo produto: %v", v.Produto))
	}

	c.JSON(http.StatusOK, gin.H{
		"produtos": pesRes,
		"msg":      recomendacao,
	})
}

func GetInsightProduto(c *gin.Context) {
	con, err := db.Conn()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}
	defer con.Close()
	id := c.Param("id")
	prod, err := db.GetProdutosInsight(con, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}
	var recomendacao []string
	for i, v := range prod {
		recomendacao = append(recomendacao, fmt.Sprintf("O funcionario %v é o %vº recomendado para essa tarefa.", v.NomeFunc, strconv.Itoa(i+1)))
	}

	c.JSON(http.StatusOK, gin.H{
		"feedback": prod,
		"msg":      recomendacao,
	})
}
