package contoroller

import (
	"godb/usecase"

	"github.com/gin-gonic/gin"
)

type ProductContoller struct {
	productUseCase usecase.ProductUseCase
}

func (pc *ProductContoller) GetAllProduct(c *gin.Context) {
	products, err := pc.productUseCase.GetAllProduct()
	if err != nil {
		c.JSON(500,gin.H{
			"message" : err.
		})
	}
}