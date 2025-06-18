package cmd

import (
	"fmt"

	"github.com/evaizee/seat-arrangements/backend/di"
	"github.com/evaizee/seat-arrangements/backend/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the HTTP server",
	Long:  `Start the HTTP server for the airplane seat booking system.`,
	Run: func(cmd *cobra.Command, args []string) {
		serve()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Add flags specific to the serve command
	serveCmd.Flags().IntP("port", "p", 8080, "Port to run the server on")
	serveCmd.Flags().StringP("host", "H", "0.0.0.0", "Host to run the server on")

	// Bind flags to viper
	viper.BindPFlag("server.port", serveCmd.Flags().Lookup("port"))
	viper.BindPFlag("server.host", serveCmd.Flags().Lookup("host"))
}

func serve() {
	// Create a new Fiber app
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			// Default 500 statuscode
			code := fiber.StatusInternalServerError

			if e, ok := err.(*fiber.Error); ok {
				// Override status code if fiber.Error type
				code = e.Code
			}

			// Set Content-Type: application/json
			c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

			// Return statuscode with error message
			return c.Status(code).JSON(fiber.Map{
				"error": true,
				"msg":   err.Error(),
			})
		},
	})

	// Use middlewares
	app.Use(recover.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: viper.GetString("cors.origins"),
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, POST, PUT, DELETE, OPTIONS",
	}))

	// Initialize dependency injection container
	container := di.NewContainer()

	// Setup routes
	routes.SetupRoutes(app, container)

	// Get host and port from config
	host := viper.GetString("server.host")
	port := viper.GetInt("server.port")

	// Start server
	serverAddr := fmt.Sprintf("%s:%d", host, port)
	zap.L().Info("Starting server", zap.String("address", serverAddr))

	if err := app.Listen(serverAddr); err != nil {
		zap.L().Fatal("Error starting server", zap.Error(err))
	}
}