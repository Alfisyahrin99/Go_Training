package delivery

import (
	"godb/config"
	"godb/manager"

	"github.com/gin-gonic/gin"
)

type appServer struct {
	engine         *gin.Engine
	useCaseManager manager.UseCaseManager
}

func Server() *appServer {
	ginEngine := gin.Default()
	config := config.NewConfig()
	infra := manager.NewInfraManager(config)
	repo := manager.NewRepositoryManager(infra)
	usecase := manager.NewUseCaseManager(repo)
	return &appServer{
		engine:         ginEngine,
		useCaseManager: usecase,
	}
}

func (a *appServer) initHandlers() {
	// controller.newProductController(a.engine, a.useCaseManager.ProductUseCase())
	// controller.NewProductController(a.engine, a.useCaseManager.ProductUseCase())
}

func (a *appServer) Run() {
	a.initHandlers()
	err := a.engine.Run(":8085")
	if err != nil {
		panic(err.Error())
	}
}