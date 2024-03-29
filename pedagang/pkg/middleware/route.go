package middleware

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nanwp/jajan-yuk/pedagang/client/product_client"
	"github.com/nanwp/jajan-yuk/pedagang/config"
	"github.com/nanwp/jajan-yuk/pedagang/core/module"
	"github.com/nanwp/jajan-yuk/pedagang/handler/api"
	"github.com/nanwp/jajan-yuk/pedagang/pkg/conn"
	"github.com/nanwp/jajan-yuk/pedagang/repository/pedagang_repository"
	"github.com/nanwp/jajan-yuk/pedagang/repository/product_repository"
	"github.com/rs/cors"
	"gorm.io/gorm"
)

func InitRouter(cfg config.Config, db *gorm.DB) (http.Handler, conn.CacheService) {
	coreRedis, _ := conn.InitRedis(cfg)

	pedagangRepository := pedagang_repository.NewPedagangRepository(cfg, db)

	productClient := product_client.NewProductClient(cfg)
	productRepo := product_repository.New(cfg, productClient)

	pedagangService := module.NewPedagangService(cfg, pedagangRepository, productRepo)
	httpHandler := api.NewHTTPHandler(pedagangService)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})

	router := mux.NewRouter()
	router.HandleFunc("/api/v1/pedagang", httpHandler.CreatePedagang).Methods("POST")
	router.HandleFunc("/api/v1/pedagang/{id}", httpHandler.GetPedagangByID).Methods("GET")
	router.HandleFunc("/api/v1/pedagang", httpHandler.GetAllPedagangNearby).Methods("GET")
	router.HandleFunc("/api/v1/pedagang/location", httpHandler.UpdateLocation).Methods("PUT")
	router.HandleFunc("/api/v1/pedagang/switch-status", httpHandler.SwitchStatus).Methods("PUT")
	router.HandleFunc("/images/{image_name}", httpHandler.GetImage).Methods("GET")
	router.HandleFunc("/api/v1/current-pedagang", httpHandler.GetCurrent).Methods("GET")

	handler := c.Handler(router)

	return handler, coreRedis
}
