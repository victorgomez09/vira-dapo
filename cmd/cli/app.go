package cli

import (
	"github.com/labstack/echo/v4"
	"github.com/urfave/cli/v2"
	"github.com/victorgomez09/vira-dapo/api/handlers"
	"github.com/victorgomez09/vira-dapo/internal/database"
)

func AppCommands() *cli.Command {
	return &cli.Command{
		Name:  "app",
		Usage: "application operations",
		Subcommands: []*cli.Command{
			{
				Name:  "start",
				Usage: "start application",
				Action: func(c *cli.Context) error {
					return startApplication()
				},
			},
		},
	}
}

func startApplication() error {
	// Create echo server
	e := echo.New()

	// Init Database
	db := &database.DB{}
	db.InitDb()

	// Initialize handlers
	authHandler := handlers.AuthHandler{
		Repository: db,
	}
	userHandler := handlers.UserHandler{
		Repository: db,
	}
	projectHandler := handlers.ProjectHandler{
		Repository: db,
	}
	collectionHandler := handlers.CollectionHandler{
		Repository: db,
	}

	// Start router with 'api' prfix
	router := e.Group("/api")

	// Auth endpoints
	authGroup := router.Group("/auth")
	authGroup.POST("/login", authHandler.Login)
	authGroup.POST("/create", userHandler.Createuser)

	// User endpoints
	// userGroup := router.Group("/users")

	// Project endpoints
	projectGroup := router.Group("/projects")
	projectGroup.GET("/", projectHandler.FindAll)
	projectGroup.GET("/:id", projectHandler.FindById)
	projectGroup.POST("/", projectHandler.Create)

	// Collection enpoints
	collectionGroup := router.Group("/collections")
	collectionGroup.POST("/", collectionHandler.Create)

	e.Logger.Fatal(e.Start(":9999"))

	return nil
}
