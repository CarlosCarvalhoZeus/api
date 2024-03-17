package handlers

import (
	"api/db"
	"net/http"

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
