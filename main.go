package main

import "github.com/gin-gonic/gin"

// import "godb/delivery"

func main() {
	// delivery.Run()

	r := gin.Default()
	r.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hallo",
		})
	})

	r.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "Selamat Sore")
	})

	err := r.Run(":8085")
	if err != nil {
		panic(err)
	}
}
