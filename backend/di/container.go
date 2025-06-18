package di

import (
	"database/sql"
	"fmt"

	"github.com/evaizee/seat-arrangements/backend/controllers"
	"github.com/evaizee/seat-arrangements/backend/repositories"
	"github.com/evaizee/seat-arrangements/backend/repositories/postgres"
	"github.com/evaizee/seat-arrangements/backend/services"
	"github.com/evaizee/seat-arrangements/backend/services/impl"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// Container holds all the dependencies for the application
type Container struct {
	// Database connection
	DB *sql.DB

	// Repositories
	UserRepository          repositories.UserRepository
	FlightRepository        repositories.FlightRepository
	AircraftRepository      repositories.AircraftRepository
	CabinRepository         repositories.CabinRepository
	SeatRepository          repositories.SeatRepository
	RowRepository          repositories.RowRepository
	BookingRepository       repositories.BookingRepository
	PassengerRepository     repositories.PassengerRepository
	FrequentFlyerRepository repositories.FrequentFlyerRepository

	// Services
	UserService          services.UserService
	FlightService        services.FlightService
	AircraftService      services.AircraftService
	CabinService         services.CabinService
	SeatService          services.SeatService
	BookingService       services.BookingService
	AuthService          services.AuthService
	PassengerService     services.PassengerService
	FrequentFlyerService services.FrequentFlyerService

	// Controllers
	// UserController     *controllers.UserController
	// FlightController   *controllers.FlightController
	// AircraftController *controllers.AircraftController
	// CabinController    *controllers.CabinController
	SeatController *controllers.SeatController
	// BookingController  *controllers.BookingController
	// AuthController     *controllers.AuthController
}

// NewContainer creates a new dependency injection container
func NewContainer() *Container {
	// Initialize the container
	container := &Container{}

	// Connect to the database
	container.initDB()

	// Initialize repositories
	container.initRepositories()

	// Initialize services
	container.initServices()

	// Initialize controllers
	container.initControllers()

	return container
}

// initDB initializes the database connection
func (c *Container) initDB() {
	// Get database configuration from viper
	dbHost := viper.GetString("database.host")
	dbPort := viper.GetInt("database.port")
	dbUser := viper.GetString("database.user")
	dbPassword := viper.GetString("database.password")
	dbName := viper.GetString("database.dbname")
	dbSSLMode := viper.GetString("database.sslmode")

	// Create the connection string
	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		dbHost, dbPort, dbUser, dbPassword, dbName, dbSSLMode,
	)

	// Connect to the database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		zap.L().Info(connStr)
		zap.L().Fatal("Failed to connect to database", zap.Error(err))
	}

	// Set connection pool settings
	db.SetMaxOpenConns(viper.GetInt("database.max_open_conns"))
	db.SetMaxIdleConns(viper.GetInt("database.max_idle_conns"))

	// Test the connection
	if err := db.Ping(); err != nil {
		zap.L().Fatal("Failed to ping database", zap.Error(err))
	}

	zap.L().Info("Connected to database")

	// Set the database connection in the container
	c.DB = db
}

// initRepositories initializes all repositories
func (c *Container) initRepositories() {
	c.UserRepository = postgres.NewUserRepository(c.DB)
	// Only uncomment these when the repository implementations are available
	c.FlightRepository = postgres.NewFlightRepository(c.DB)
	c.AircraftRepository = postgres.NewAircraftRepository(c.DB)
	c.CabinRepository = postgres.NewCabinRepository(c.DB)
	c.SeatRepository = postgres.NewSeatRepository(c.DB)
	c.RowRepository = postgres.NewRowRepository(c.DB)
	// c.BookingRepository = postgres.NewBookingRepository(c.DB)
	c.PassengerRepository = postgres.NewPassengerRepository(c.DB)
	c.FrequentFlyerRepository = postgres.NewFrequentFlyerRepository(c.DB)
}

// initServices initializes all services
func (c *Container) initServices() {

	// Initialize SeatService with just the repositories we have available
	c.SeatService = impl.NewSeatService(
		c.SeatRepository,
		c.PassengerRepository,
		c.AircraftRepository,
		c.CabinRepository,
		c.FlightRepository,
		c.RowRepository,
	)

	// c.BookingService = impl.NewBookingService(c.BookingRepository, c.SeatRepository)
	// c.AuthService = impl.NewAuthService(c.UserRepository)
	c.PassengerService = impl.NewPassengerService(c.PassengerRepository, c.FrequentFlyerRepository)
	c.FrequentFlyerService = impl.NewFrequentFlyerService(c.FrequentFlyerRepository)
}

// initControllers initializes all controllers
func (c *Container) initControllers() {
	// c.UserController = controllers.NewUserController(c.UserService)
	// c.FlightController = controllers.NewFlightController(c.FlightService)
	// c.AircraftController = controllers.NewAircraftController(c.AircraftService)
	// c.CabinController = controllers.NewCabinController(c.CabinService)
	c.SeatController = controllers.NewSeatController(c.SeatService)
	// c.BookingController = controllers.NewBookingController(c.BookingService)
	// c.AuthController = controllers.NewAuthController(c.AuthService)
}
