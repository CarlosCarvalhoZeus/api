package routes

import (
	"api/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GinSetup() {
	router := gin.Default()

	initializeRoutes(router)
}
func initializeRoutes(router *gin.Engine) {
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"Success": "API up and running...",
		})
	})
	router.GET("/pessoas", handlers.GetPessoas)
	router.GET("/produtos", handlers.GetProdutos)
	router.POST("/os", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"success": "under development...",
		})
	})
	router.GET("produtos/:id")

	// router.GET("/:name")
	router.Run()

}
