package controller

import (
	"godb/model"
	"godb/usecase"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productUseCase usecase.ProductUseCase
}

func (pc *ProductController) GetAllProduct(c *gin.Context) {
	products, err := pc.productUseCase.GetAllProduct()
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"message": "OK",
			"data":    products,
		})
	}
}

func (pc *ProductController) CreateNewProduct(c *gin.Context) {
	var newProduct *model.Product
	err := c.Bind(&newProduct)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
	} else {
		pc.productUseCase.CreateNewProduct(newProduct)
		c.JSON(200, gin.H{
			"message": "OK",
			"data":    newProduct,
		})
	}
}

func NewProductController(router *gin.Engine, productUc usecase.ProductUseCase) *ProductController {
	newProductController := ProductController{
		productUseCase: productUc,
	}

	router.GET("/products", newProductController.GetAllProduct)
	router.POST("/product", newProductController.CreateNewProduct)
	//router.GET()
	return &newProductController
}
