package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"
	"todo-list/internal/shared/databases"
	"todo-list/internal/shared/databases/migrations"

	"github.com/goccy/go-json"
	"github.com/pkg/errors"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/utils"

	"todo-list/internal/activity"
	activityController "todo-list/internal/activity/controllers"
	activityRepo "todo-list/internal/activity/repositories"
	activityService "todo-list/internal/activity/services"
	activityValidator "todo-list/internal/activity/validators"

	"todo-list/internal/todo"
	todoController "todo-list/internal/todo/controllers"
	todoRepo "todo-list/internal/todo/repositories"
	todoService "todo-list/internal/todo/services"
	todoValidator "todo-list/internal/todo/validators"
)

// validate env
func init() {
	if val, ok := os.LookupEnv("APP_PORT"); val == "" || !ok {
		log.Fatal("please provide APP_PORT environment variable")
	}

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
		Expiration: 1 * time.Minute,
		KeyGenerator: func(c *fiber.Ctx) string {
			return utils.CopyString(c.OriginalURL())
		},
	}))

	// INITIALIZE DATABASES
	db, err := databases.MySQLConn()
	if err != nil {
		log.Fatal("error during opening mysql client", err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Print("error during closing mysql client", errors.WithStack(err))
		}
	}(db)

	// MIGRATTIONS
	err = migrations.Migrate(db)
	if err != nil {
		log.Fatal("error during migrating database", errors.WithStack(err))
	}

	// INITIALIZE REPOSITORIES
	activityRepoImpl := activityRepo.NewActivity(db)
	todoRepoImpl := todoRepo.NewTodo(db)

	// INITIALIZE VALIDATORS
	activityValidatorImpl := activityValidator.NewActivity()
	todoValidatorImpl := todoValidator.NewTodo()

	// INITIALIZE SERVICES
	activityServiceImpl := activityService.NewActivity(
		activityValidatorImpl,
		activityRepoImpl,
	)
	todoServiceImpl := todoService.NewTodo(todoValidatorImpl, todoRepoImpl)

	// INITIALIZE CONTROLLERS
	activityControllerImpl := activityController.NewActivity(activityServiceImpl)
	todoControllerImpl := todoController.NewTodo(todoServiceImpl)

	// INTIALIZE ROUTES
	activity.NewRoute(app, activityControllerImpl)
	todo.NewRoute(app, todoControllerImpl)

	app.Listen(fmt.Sprintf(":%s", os.Getenv("APP_PORT")))
}
