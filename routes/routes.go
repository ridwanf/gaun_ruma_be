package routes

import (
	"gaunRumaRestApi/cmd/handler"
	"gaunRumaRestApi/config"
	"gaunRumaRestApi/config/db"
	"gaunRumaRestApi/repository"
	"gaunRumaRestApi/services"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	echojwt "github.com/labstack/echo-jwt/v4"

	"github.com/labstack/echo/v4"
)

func Init(d db.Handler, conf config.Configuration) *echo.Echo {
	userRepository := repository.NewUserRepository(&d)
	productTypeRepository := repository.NewProductTypeRepository(&d)

	// configure middleware with the custom claims type
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(handler.JwtCustomClaims)
		},
		SigningKey: []byte(conf.SIGNING_KEY),
	}

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello this is gaun ruma")
	})

	productType := e.Group("/product-type")

	productType.Use(echojwt.WithConfig(config))
	handlerProductType := handler.NewProductTypeHandler(productTypeRepository)
	{
		productType.GET("", handlerProductType.GetAllProductType)
		productType.POST("", handlerProductType.CreateProductType)
		productType.PUT("", handlerProductType.UpdateProductType)
		productType.DELETE("", handlerProductType.DeleteProductType)
	}

	stockApi := e.Group("/stock")

	stockApi.Use(echojwt.WithConfig(config))
	handlerStock := handler.NewProductHandler(services.NewProductService(repository.NewProductRepository(&d)))
	{
		stockApi.GET("", handlerStock.GetAllProduct)
		stockApi.GET("/:id", handlerStock.GetById)
		stockApi.POST("", handlerStock.CreateProduct)
		stockApi.PUT("", handlerStock.UpdateProduct)
		stockApi.DELETE("", handlerStock.DeleteProduct)
	}

	e.POST("/login", handler.NewUserHandler(userRepository).Login)
	e.POST("/generate-password", handler.NewUserHandler(userRepository).GenerateHasPassword)
	e.POST("/register", handler.NewUserHandler(userRepository).Register)

	return e
}
