package main

import (
	"log"
	"os"
	"time"
	"todo-list/internal/activity"
	"todo-list/internal/shared/databases"

	"github.com/goccy/go-json"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"

	activityController "todo-list/internal/activity/controllers"
	activityRepo "todo-list/internal/activity/repositories"
	activityService "todo-list/internal/activity/services"
	activityValidator "todo-list/internal/activity/validators"
)

// validate env
func init() {
	if val, ok := os.LookupEnv("MYSQL_USER"); val == "" || !ok {
		log.Fatal("please provide MYSQL_USER enviroment variable")
	}

	if _, ok := os.LookupEnv("MYSQL_PASSWORD"); !ok {
		log.Fatal("please provide MYSQL_PASSWORD environment variable")
	}

	if val, ok := os.LookupEnv("MYSQL_HOST"); val == "" || !ok {
		log.Fatal("please provide MYSQL_HOST environment variable")
	}

	if val, ok := os.LookupEnv("MYSQL_PORT"); val == "" || !ok {
		log.Fatal("please provide MYSQL_PORT environment variable")
	}

	if val, ok := os.LookupEnv("MYSQL_DBNAME"); val == "" || !ok {
		log.Fatal("please provide MYSQL_DBNAME environment variable")
	}
}

func main() {
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})
	app.Use(cache.New(cache.Config{
		Expiration: 2 * time.Minute,
	}))

	// INITIALIZE DATABASES
	db, err := databases.MySQLConn()
	if err != nil {
		log.Fatal("error during connecting to database", err)
	}

	// INITIALIZE REPOSITORIES
	activityRepoImpl := activityRepo.NewActivity(db)

	// INITIALIZE VALIDATORS
	activityValidatorImpl := activityValidator.NewActivity()

	// INITIALIZE SERVICES
	activityServiceImpl := activityService.NewActivity(
		activityValidatorImpl,
		activityRepoImpl,
	)

	// INITIALIZE CONTROLLERS
	activityControllerImpl := activityController.NewActivity(activityServiceImpl)

	// INTIALIZE ROUTES
	activity.NewRoute(app, activityControllerImpl)

	app.Listen(":3030")
}
