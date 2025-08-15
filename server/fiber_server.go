package server

import (
	"fmt"

	"github.com/HanThamarat/GO-Bucket-Service/packages/conf"
	_ "github.com/HanThamarat/GO-Bucket-Service/packages/doc"
	"github.com/HanThamarat/GO-Bucket-Service/packages/hooks"
	"github.com/HanThamarat/GO-Bucket-Service/packages/middlewares"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
)

type (
	Server interface {
		Start()
	}

	fiberServer struct {
		app *fiber.App
		conf *conf.Config
	}
)

func NewFiberServer(conf *conf.Config) Server {
	fiberApp := fiber.New(fiber.Config{
		ReadBufferSize: 60 * 1024,
		DisableStartupMessage: false,
		AppName: conf.App.NAME,
	});

		return &fiberServer{
			app: fiberApp,
			conf: conf,
		}
}

func (s *fiberServer) Start() {
	s.app.Use(logger.New());

	s.app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Authorization, Content-Type",
	}));

	api := s.app.Group("/api", hooks.DecryptJWT);

	api.Use(middlewares.AuthMiddleware())
	api.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(s.conf.JWT.Secret)},
	}))

	s.app.Get("", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "api is available",
		})
	});

	s.app.Get("/swagger/*", swagger.HandlerDefault);
	s.InitializeFileService(api);
		
	serverUrl := fmt.Sprintf(":%d", s.conf.Server.Port);


	s.app.Listen(serverUrl);
}