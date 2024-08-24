package server

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/unawaretub86/leal-challenge/cmd/server/dependencies"
	configuration "github.com/unawaretub86/leal-challenge/config"
	"github.com/unawaretub86/leal-challenge/src/adapters/http"
	"github.com/unawaretub86/leal-challenge/src/adapters/repository"
	"github.com/unawaretub86/leal-challenge/src/domain/ports"
	"github.com/unawaretub86/leal-challenge/src/domain/services"
)

type Service struct {
	leal ports.LealPort
}

type Initiator struct {
	service Service
	router  *gin.Engine
}

func NewInitiator() *Initiator {
	return &Initiator{
		service: Service{},
		router:  &gin.Engine{},
	}
}


func (i *Initiator) initService() {
	lealRepo := repository.NewLealRepository(i.initDB())
	lealService := services.NewLealService(lealRepo)

	i.service = Service{
		lealService,
	}
}

func (i *Initiator) initRouter() {
	i.router = gin.Default()

	lealRouter := http.NewRouter(i.service.leal)
	lealRouter.SetRoutes(i.router)
}

func (i *Initiator) GetRouter() *gin.Engine {
	return i.router
}

func (i *Initiator) initDB() *gorm.DB {
	dbUrl := configuration.GetDatabaseConfig()

	db, err := dependencies.ConnectDB(dbUrl)
	if err != nil {
		panic(err.Error())
	}

	return db.Connection()
}

func (i *Initiator) InitAll() {
	i.initDB()
	i.initRouter()
	i.initService()

	err := i.GetRouter().Run(":" + "8080")
	if err != nil {
		log.Fatal("Server init error:", err.Error())
	}
}
