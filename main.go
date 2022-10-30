package main

import (
	"mypdf/bill"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	engine.GET("/order/bile", func(ctx *gin.Context) {
		bilePDF, err := bill.GeneratePDF()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		ctx.Data(http.StatusOK, "application/pdf", bilePDF.Bytes())
	})
	engine.Run(":3000")
}
